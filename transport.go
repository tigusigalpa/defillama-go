package defillama

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"syscall"
	"time"
)

// transport executes registered GET routes. All services share one instance.
type transport struct {
	cfg    config
	client *http.Client
}

// responseDrainLimit bounds best-effort cleanup after the configured response
// limit or a read failure. It is separate from retained response data.
const responseDrainLimit int64 = 1 << 20

func newTransport(cfg config) *transport {
	return &transport{cfg: cfg, client: cfg.http()}
}

// get resolves the route, performs the request (with retries) and decodes the
// JSON body into out.
func (t *transport) get(ctx context.Context, routeID string, pathParams map[string]any, query url.Values, out any) error {
	r, err := lookupRoute(routeID)
	if err != nil {
		return err
	}
	if err := validateQuery(r, query); err != nil {
		return err
	}
	u, redacted, err := t.resolveFor(r, pathParams)
	if err != nil {
		return err
	}
	if len(query) > 0 {
		qs := query.Encode()
		u += "?" + qs
		redacted += "?" + qs
	}
	keyInURL := r.Tier == "pro" || (t.cfg.preferProForFree && t.cfg.apiKey != "" && r.ProPath != "")
	_, err = t.doWithRetry(ctx, r, u, redacted, keyInURL, out)
	return err
}

// getReceipt performs a registered request and returns the receipt for its
// final response. A configured ReceiptObserver still receives every response,
// including responses that trigger a retry.
func (t *transport) getReceipt(ctx context.Context, routeID string, pathParams map[string]any, query url.Values) (*ResponseReceipt, error) {
	r, err := lookupRoute(routeID)
	if err != nil {
		return nil, err
	}
	if err := validateQuery(r, query); err != nil {
		return nil, err
	}
	u, redacted, err := t.resolveFor(r, pathParams)
	if err != nil {
		return nil, err
	}
	if len(query) > 0 {
		qs := query.Encode()
		u += "?" + qs
		redacted += "?" + qs
	}
	keyInURL := r.Tier == "pro" || (t.cfg.preferProForFree && t.cfg.apiKey != "" && r.ProPath != "")
	return t.doWithRetry(ctx, r, u, redacted, keyInURL, nil)
}

// resolveFor delegates URL building to the shared config resolver.
func (t *transport) resolveFor(r route, pathParams map[string]any) (string, string, error) {
	return t.cfg.resolve(r, pathParams)
}

func (t *transport) doWithRetry(ctx context.Context, r route, u, redacted string, keyInURL bool, out any) (*ResponseReceipt, error) {
	for attempt := 1; ; attempt++ {
		receipt, err := t.doOnce(ctx, r, u, redacted, keyInURL, attempt, out)
		if err == nil {
			return receipt, nil
		}
		if ctx.Err() != nil {
			return receipt, err
		}
		if !t.retryable(err, attempt) {
			return receipt, err
		}
		delay := t.cfg.retry.delay(attempt)
		var rl *RateLimitError
		if errors.As(err, &rl) && rl.RetryAfter > 0 && rl.RetryAfter > delay {
			delay = rl.RetryAfter
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return receipt, &TransportError{URL: redacted, Err: ctx.Err()}
		case <-timer.C:
		}
	}
}

// retryable reports whether a failed attempt may be retried.
func (t *transport) retryable(err error, attempt int) bool {
	if attempt >= t.cfg.retry.MaxAttempts {
		return false
	}
	if errors.Is(err, ErrResponseBodyTooLarge) || errors.Is(err, context.Canceled) {
		return false
	}
	var transportErr *TransportError
	if errors.As(err, &transportErr) && isTransientTransportError(transportErr.Err) {
		return true
	}
	var decodeErr *DecodeError
	if errors.As(err, &decodeErr) && isTransientTransportError(decodeErr.Err) {
		return true
	}
	var rateLimitErr *RateLimitError
	if errors.As(err, &rateLimitErr) {
		return true
	}
	var apiErr *APIError
	if errors.As(err, &apiErr) {
		return apiErr.StatusCode >= 500 && apiErr.StatusCode <= 599
	}
	return false
}

func isTransientTransportError(err error) bool {
	if errors.Is(err, context.Canceled) {
		return false
	}
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) && (dnsErr.IsTemporary || dnsErr.IsTimeout) {
		return true
	}
	for cause := err; cause != nil; cause = errors.Unwrap(cause) {
		if netErr, ok := cause.(net.Error); ok && netErr.Timeout() {
			return true
		}
	}
	return errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) ||
		errors.Is(err, syscall.ECONNRESET) || errors.Is(err, syscall.ECONNREFUSED) ||
		errors.Is(err, syscall.EPIPE)
}

func (t *transport) doOnce(ctx context.Context, r route, u, redacted string, keyInURL bool, attempt int, out any) (*ResponseReceipt, error) {
	apiKey := ""
	if keyInURL {
		apiKey = t.cfg.apiKey
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, &TransportError{URL: redacted, Err: redactTransportError(err, redacted, apiKey)}
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", t.cfg.userAgentHeader())

	client := t.client
	if keyInURL {
		// net/http follows redirects by default. A Pro request has its key in
		// the URL, so never follow a redirect to another origin.
		redirectClient := *client
		previousCheck := redirectClient.CheckRedirect
		redirectClient.CheckRedirect = func(next *http.Request, via []*http.Request) error {
			if len(via) > 0 && !sameOrigin(next.URL, via[0].URL) {
				return http.ErrUseLastResponse
			}
			if previousCheck != nil {
				return previousCheck(next, via)
			}
			if len(via) >= 10 {
				return errors.New("stopped after 10 redirects")
			}
			return nil
		}
		client = &redirectClient
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, &TransportError{URL: redacted, Err: redactTransportError(err, redacted, apiKey)}
	}
	capturedAt := time.Now().UTC()
	originalBody := resp.Body
	body := finalizeResponseBody(originalBody, t.cfg.maxResponseBytes)
	capture := out == nil || t.cfg.receiptObserver != nil

	var primary error
	readIsPrimary := false
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		if body.readErr != nil {
			primary = &DecodeError{URL: redacted, Err: body.readErr}
			readIsPrimary = true
		} else if out != nil {
			if err := decodeResponseJSON(bytes.NewReader(body.data), out); err != nil {
				primary = &DecodeError{URL: redacted, Err: err}
			}
		}
	} else {
		primary = apiErrorFromResponse(resp.StatusCode, resp.Header, body.data, redacted, apiKey)
	}
	finalErr := joinResponseErrors(primary, redacted, body, readIsPrimary)

	var receipt *ResponseReceipt
	if capture {
		captured := newResponseReceipt(
			r.Method+" "+r.Path,
			redacted,
			r.DocsURL,
			resp.StatusCode,
			redactHeader(resp.Header, apiKey),
			body.data,
			attempt,
			capturedAt,
			body.completedAt,
			body.complete,
		)
		receipt = &captured
		if t.cfg.receiptObserver != nil {
			t.cfg.receiptObserver(captured)
		}
	}
	return receipt, finalErr
}

type finalizedResponseBody struct {
	data        []byte
	readErr     error
	drainErr    error
	closeErr    error
	complete    bool
	completedAt time.Time
}

// finalizeResponseBody reads at most max bytes for retention, then drains and
// closes the original response body exactly once. It never receives a replay
// buffer, so cleanup cannot be redirected to a replacement body.
func finalizeResponseBody(original io.ReadCloser, limit int64) finalizedResponseBody {
	data, reachedEOF, readErr := readResponseBody(original, limit)
	_, drainErr := io.Copy(io.Discard, io.LimitReader(original, responseDrainLimit))
	closeErr := original.Close()
	return finalizedResponseBody{
		data:        data,
		readErr:     readErr,
		drainErr:    drainErr,
		closeErr:    closeErr,
		complete:    reachedEOF && drainErr == nil && closeErr == nil,
		completedAt: time.Now().UTC(),
	}
}

func readResponseBody(body io.Reader, limit int64) ([]byte, bool, error) {
	data, err := io.ReadAll(io.LimitReader(body, limit+1))
	if int64(len(data)) > limit {
		return data[:limit], false, &ResponseBodyTooLargeError{Limit: limit}
	}
	if err != nil {
		return data, false, err
	}
	return data, true, nil
}

func joinResponseErrors(primary error, redactedURL string, body finalizedResponseBody, readIsPrimary bool) error {
	errs := make([]error, 0, 4)
	if primary != nil {
		errs = append(errs, primary)
	}
	if body.readErr != nil && !readIsPrimary {
		errs = append(errs, &TransportError{URL: redactedURL, Err: body.readErr})
	}
	if body.drainErr != nil {
		errs = append(errs, &TransportError{URL: redactedURL, Err: body.drainErr})
	}
	if body.closeErr != nil {
		errs = append(errs, &TransportError{URL: redactedURL, Err: body.closeErr})
	}
	switch len(errs) {
	case 0:
		return nil
	case 1:
		return errs[0]
	default:
		return errors.Join(errs...)
	}
}

func decodeResponseJSON(body io.Reader, out any) error {
	dec := json.NewDecoder(body)
	if err := dec.Decode(out); err != nil {
		return err
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		if err == nil {
			return errors.New("multiple JSON values in response")
		}
		return err
	}
	return nil
}

func sameOrigin(a, b *url.URL) bool {
	return strings.EqualFold(a.Scheme, b.Scheme) && strings.EqualFold(a.Host, b.Host)
}

// apiErrorFromResponse maps a non-2xx response and its bounded body prefix to
// the typed error tree. Lifecycle errors are joined by the caller.
func apiErrorFromResponse(statusCode int, header http.Header, responseBody []byte, redactedURL, apiKey string) error {
	body := responseBody
	if len(body) > apiErrorBodyLimit {
		body = body[:apiErrorBodyLimit]
	}
	base := &APIError{
		StatusCode: statusCode,
		URL:        redactedURL,
		Header:     redactHeader(header, apiKey),
		Body:       redactBytes(body, apiKey),
	}
	switch statusCode {
	case http.StatusNotFound:
		return &NotFoundError{APIError: base}
	case http.StatusTooManyRequests:
		return &RateLimitError{APIError: base, RetryAfter: parseRetryAfter(header.Get("Retry-After"))}
	default:
		return base
	}
}

// redactTransportError removes the request URL from url.Errors returned by
// net/http and sanitizes nested error messages. The original causes remain
// available through errors.Is/As.
func redactTransportError(err error, redactedURL, apiKey string) error {
	urlErr, ok := err.(*url.Error)
	if !ok {
		return &redactedTransportCause{err: err, message: redactString(err.Error(), apiKey)}
	}
	clone := *urlErr
	clone.URL = redactedURL
	clone.Err = redactTransportError(urlErr.Err, redactedURL, apiKey)
	return &clone
}

// parseRetryAfter parses a Retry-After header (delay-seconds or HTTP-date).
func parseRetryAfter(v string) time.Duration {
	if v == "" {
		return 0
	}
	if secs, err := strconv.ParseFloat(v, 64); err == nil {
		return time.Duration(secs * float64(time.Second))
	}
	if ts, err := http.ParseTime(v); err == nil {
		if d := time.Until(ts); d > 0 {
			return d
		}
	}
	return 0
}

// rand63n is a small indirection for tests.
var rand63n = func(n int64) int64 {
	return rand.Int63n(n)
}
