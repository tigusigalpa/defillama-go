package defillama

import (
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
	return t.doWithRetry(ctx, u, redacted, keyInURL, out)
}

// resolveFor delegates URL building to the shared config resolver.
func (t *transport) resolveFor(r route, pathParams map[string]any) (string, string, error) {
	return t.cfg.resolve(r, pathParams)
}

func (t *transport) doWithRetry(ctx context.Context, u, redacted string, keyInURL bool, out any) error {
	for attempt := 1; ; attempt++ {
		err := t.doOnce(ctx, u, redacted, keyInURL, out)
		if err == nil {
			return nil
		}
		if ctx.Err() != nil {
			return err
		}
		if !t.retryable(err, attempt) {
			return err
		}
		delay := t.cfg.retry.delay(attempt)
		if rl, ok := err.(*RateLimitError); ok && rl.RetryAfter > 0 && rl.RetryAfter > delay {
			delay = rl.RetryAfter
		}
		timer := time.NewTimer(delay)
		select {
		case <-ctx.Done():
			timer.Stop()
			return &TransportError{URL: redacted, Err: ctx.Err()}
		case <-timer.C:
		}
	}
}

// retryable reports whether a failed attempt may be retried.
func (t *transport) retryable(err error, attempt int) bool {
	if attempt >= t.cfg.retry.MaxAttempts {
		return false
	}
	switch e := err.(type) {
	case *TransportError:
		return isTransientTransportError(e.Err)
	case *RateLimitError:
		return true
	case *APIError:
		return e.StatusCode >= 500 && e.StatusCode <= 599
	default:
		return false
	}
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

func (t *transport) doOnce(ctx context.Context, u, redacted string, keyInURL bool, out any) error {
	apiKey := ""
	if keyInURL {
		apiKey = t.cfg.apiKey
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return &TransportError{URL: redacted, Err: redactTransportError(err, redacted, apiKey)}
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", t.cfg.userAgentHeader())

	client := t.client
	if keyInURL {
		// net/http follows redirects by default. A Pro request has its key in
		// the URL, so never follow a redirect to another origin.
		copy := *client
		previousCheck := copy.CheckRedirect
		copy.CheckRedirect = func(next *http.Request, via []*http.Request) error {
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
		client = &copy
	}

	resp, err := client.Do(req)
	if err != nil {
		return &TransportError{URL: redacted, Err: redactTransportError(err, redacted, apiKey)}
	}
	defer func() {
		io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<20))
		resp.Body.Close()
	}()

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		dec := json.NewDecoder(resp.Body)
		if err := dec.Decode(out); err != nil {
			return &DecodeError{URL: redacted, Err: err}
		}
		var extra any
		if err := dec.Decode(&extra); err != io.EOF {
			if err == nil {
				err = errors.New("multiple JSON values in response")
			}
			return &DecodeError{URL: redacted, Err: err}
		}
		return nil
	}
	return apiErrorFromResponse(resp, redacted, apiKey)
}

func sameOrigin(a, b *url.URL) bool {
	return strings.EqualFold(a.Scheme, b.Scheme) && strings.EqualFold(a.Host, b.Host)
}

// apiErrorFromResponse maps a non-2xx response to the typed error tree.
func apiErrorFromResponse(resp *http.Response, redactedURL, apiKey string) error {
	body, _ := io.ReadAll(io.LimitReader(resp.Body, apiErrorBodyLimit))
	base := &APIError{
		StatusCode: resp.StatusCode,
		URL:        redactedURL,
		Header:     redactHeader(resp.Header, apiKey),
		Body:       redactBytes(body, apiKey),
	}
	switch {
	case resp.StatusCode == http.StatusNotFound:
		return &NotFoundError{APIError: base}
	case resp.StatusCode == http.StatusTooManyRequests:
		return &RateLimitError{APIError: base, RetryAfter: parseRetryAfter(resp.Header.Get("Retry-After"))}
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
