package defillama

import (
	"bytes"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
	"sync"
	"testing"
	"time"
)

func TestResponseBodyLifecycleClosesOriginalExactlyOnce(t *testing.T) {
	valid := []byte(`[{"id":"provider-aave","tvl":1,"chainTvls":{"Ethereum":1}}]`)
	tests := []struct {
		name     string
		options  []Option
		call     func(*Client) (*ResponseReceipt, error)
		wantBody bool
	}{
		{
			name: "receipt call",
			call: func(c *Client) (*ResponseReceipt, error) {
				return c.TVL().GetProtocolsReceipt(context.Background())
			},
			wantBody: true,
		},
		{
			name:    "typed call with observer",
			options: []Option{WithReceiptObserver(func(ResponseReceipt) {})},
			call: func(c *Client) (*ResponseReceipt, error) {
				_, err := c.TVL().GetProtocols(context.Background())
				return nil, err
			},
		},
		{
			name: "typed call without capture",
			call: func(c *Client) (*ResponseReceipt, error) {
				_, err := c.TVL().GetProtocols(context.Background())
				return nil, err
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body := newTrackedReadCloser(valid, nil)
			client := lifecycleClient(t, http.StatusOK, body, tt.options...)
			receipt, err := tt.call(client)
			if err != nil {
				t.Fatal(err)
			}
			if got := body.closeCount(); got != 1 {
				t.Errorf("original body Close calls = %d, want 1", got)
			}
			if tt.wantBody && (receipt == nil || !receipt.Complete) {
				t.Errorf("receipt = %#v, want a complete receipt", receipt)
			}
		})
	}
}

func TestReceiptFinalizesBeforeObserverAndJoinsCloseError(t *testing.T) {
	closeSentinel := errors.New("close sentinel")
	bodyBytes := []byte(`[{"id":"provider-aave","tvl":1,"chainTvls":{"Ethereum":1}}]`)
	body := newTrackedReadCloser(bodyBytes, closeSentinel)
	var observed []ResponseReceipt
	var closedAtObservation bool
	client := lifecycleClient(t, http.StatusOK, body,
		WithReceiptObserver(func(receipt ResponseReceipt) {
			observed = append(observed, receipt)
			closedAtObservation = body.closeCount() == 1
		}),
	)

	receipt, err := client.TVL().GetProtocolsReceipt(context.Background())
	if !errors.Is(err, closeSentinel) {
		t.Fatalf("err = %v, want close sentinel", err)
	}
	var transportErr *TransportError
	if !errors.As(err, &transportErr) {
		t.Fatalf("err = %v, want TransportError", err)
	}
	if receipt == nil || receipt.Complete {
		t.Fatalf("receipt = %#v, want incomplete evidence", receipt)
	}
	if receipt.CompletedAt.Before(receipt.CapturedAt) {
		t.Errorf("CompletedAt %s before CapturedAt %s", receipt.CompletedAt, receipt.CapturedAt)
	}
	if got := receipt.BodySHA256(); got != sha256Hex(bodyBytes) {
		t.Errorf("partial/full body hash = %s", got)
	}
	if len(observed) != 1 || !closedAtObservation {
		t.Errorf("observer calls = %d, closed at observation = %v", len(observed), closedAtObservation)
	}
	if got := body.closeCount(); got != 1 {
		t.Errorf("original body Close calls = %d, want 1", got)
	}
}

func TestReceiptRetainsReadCancellationCause(t *testing.T) {
	body := newPartialReadCloser([]byte(`[`), context.Canceled, nil)
	client := lifecycleClient(t, http.StatusOK, body)

	receipt, err := client.TVL().GetProtocolsReceipt(context.Background())
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("err = %v, want context.Canceled", err)
	}
	var decodeErr *DecodeError
	if !errors.As(err, &decodeErr) {
		t.Fatalf("err = %v, want DecodeError", err)
	}
	if receipt == nil || receipt.Complete || string(receipt.Body()) != "[" {
		t.Fatalf("receipt = %#v, want incomplete partial evidence", receipt)
	}
	if got := body.closeCount(); got != 1 {
		t.Errorf("original body Close calls = %d, want 1", got)
	}
}

func TestReceiptRetainsPartialEvidenceAndJoinedAPIErrors(t *testing.T) {
	readSentinel := errors.New("partial read sentinel")
	closeSentinel := errors.New("close sentinel")
	partial := []byte(`{"error":"rate`)
	body := newPartialReadCloser(partial, readSentinel, closeSentinel)
	client := lifecycleClient(t, http.StatusTooManyRequests, body)

	receipt, err := client.TVL().GetProtocolsReceipt(context.Background())
	if receipt == nil || receipt.Complete {
		t.Fatalf("receipt = %#v, want incomplete partial evidence", receipt)
	}
	if !bytes.Equal(receipt.Body(), partial) || receipt.BodySHA256() != sha256Hex(partial) {
		t.Error("partial receipt body or hash was not retained")
	}
	if !errors.Is(err, readSentinel) || !errors.Is(err, closeSentinel) {
		t.Fatalf("err = %v, want both read and close causes", err)
	}
	var rateLimit *RateLimitError
	if !errors.As(err, &rateLimit) {
		t.Fatalf("err = %v, want RateLimitError", err)
	}
	if rateLimit.StatusCode != http.StatusTooManyRequests || !bytes.Equal(rateLimit.Body, partial) {
		t.Errorf("rate-limit error = status %d body %q", rateLimit.StatusCode, rateLimit.Body)
	}
	if got := body.closeCount(); got != 1 {
		t.Errorf("original body Close calls = %d, want 1", got)
	}
}

func TestTransportFailureProducesNoReceiptOrObserverEvent(t *testing.T) {
	sentinel := errors.New("transport sentinel")
	observerCalls := 0
	client, err := New(
		WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			return nil, sentinel
		})}),
		WithReceiptObserver(func(ResponseReceipt) {
			observerCalls++
		}),
	)
	if err != nil {
		t.Fatal(err)
	}

	receipt, err := client.TVL().GetProtocolsReceipt(context.Background())
	if !errors.Is(err, sentinel) {
		t.Fatalf("err = %v, want transport sentinel", err)
	}
	if receipt != nil {
		t.Errorf("receipt = %#v, want nil", receipt)
	}
	if observerCalls != 0 {
		t.Errorf("observer calls = %d, want 0", observerCalls)
	}
}

func TestResponseBodyLimit(t *testing.T) {
	t.Run("exact max plus EOF is complete", func(t *testing.T) {
		body := newTrackedReadCloser([]byte("null"), nil)
		client := lifecycleClient(t, http.StatusOK, body, WithMaxResponseBodyBytes(4))
		receipt, err := client.TVL().GetProtocolsReceipt(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if receipt == nil || !receipt.Complete || string(receipt.Body()) != "null" {
			t.Fatalf("receipt = %#v", receipt)
		}
		if got := body.closeCount(); got != 1 {
			t.Errorf("original body Close calls = %d, want 1", got)
		}
	})

	t.Run("capture overflow retains prefix", func(t *testing.T) {
		body := newTrackedReadCloser([]byte("12345"), nil)
		client := lifecycleClient(t, http.StatusOK, body, WithMaxResponseBodyBytes(4))
		receipt, err := client.TVL().GetProtocolsReceipt(context.Background())
		assertResponseBodyTooLarge(t, err, 4)
		if receipt == nil || receipt.Complete || string(receipt.Body()) != "1234" {
			t.Fatalf("receipt = %#v", receipt)
		}
		if got := body.closeCount(); got != 1 {
			t.Errorf("original body Close calls = %d, want 1", got)
		}
	})

	t.Run("typed no-capture overflow has no receipt", func(t *testing.T) {
		body := newTrackedReadCloser([]byte("[{}]"), nil)
		client := lifecycleClient(t, http.StatusOK, body, WithMaxResponseBodyBytes(3))
		route, err := lookupRoute("GET /protocols")
		if err != nil {
			t.Fatal(err)
		}
		var out []Protocol
		receipt, err := client.t.doWithRetry(context.Background(), route, "https://api.llama.fi/protocols", "https://api.llama.fi/protocols", false, &out)
		assertResponseBodyTooLarge(t, err, 3)
		if receipt != nil {
			t.Errorf("typed no-capture returned fabricated receipt: %#v", receipt)
		}
		if got := body.closeCount(); got != 1 {
			t.Errorf("original body Close calls = %d, want 1", got)
		}
	})

	t.Run("overflow is not retried", func(t *testing.T) {
		attempts := 0
		client, err := New(
			WithHTTPClient(&http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				attempts++
				return &http.Response{
					StatusCode: http.StatusOK,
					Header:     make(http.Header),
					Body:       newTrackedReadCloser([]byte("[{}]"), nil),
					Request:    r,
				}, nil
			})}),
			WithMaxResponseBodyBytes(3),
			WithRetryPolicy(RetryPolicy{MaxAttempts: 2, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}),
		)
		if err != nil {
			t.Fatal(err)
		}
		_, err = client.TVL().GetProtocols(context.Background())
		assertResponseBodyTooLarge(t, err, 3)
		if attempts != 1 {
			t.Errorf("attempts = %d, want 1", attempts)
		}
	})
}

func TestReceiptObserverFinalizesRetryAttemptsBeforeNextAttempt(t *testing.T) {
	first := newTrackedReadCloser([]byte(`{"error":"slow down"}`), nil)
	second := newTrackedReadCloser([]byte(`[{"id":"provider-aave","tvl":1,"chainTvls":{"Ethereum":1}}]`), nil)
	bodies := []*trackedReadCloser{first, second}
	var requests int
	var observed []ResponseReceipt
	var closeCountsAtObservation []int
	client, err := New(
		WithHTTPClient(&http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			body := bodies[requests]
			status := http.StatusTooManyRequests
			if requests == 1 {
				status = http.StatusOK
			}
			requests++
			return &http.Response{StatusCode: status, Header: make(http.Header), Body: body, Request: r}, nil
		})}),
		WithRetryPolicy(RetryPolicy{MaxAttempts: 2, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}),
		WithReceiptObserver(func(receipt ResponseReceipt) {
			observed = append(observed, receipt)
			closeCountsAtObservation = append(closeCountsAtObservation, bodies[receipt.Attempt-1].closeCount())
		}),
	)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := client.TVL().GetProtocols(context.Background()); err != nil {
		t.Fatal(err)
	}
	if requests != 2 || len(observed) != 2 {
		t.Fatalf("requests/observations = %d/%d, want 2/2", requests, len(observed))
	}
	for i, receipt := range observed {
		if receipt.Attempt != i+1 || !receipt.Complete || closeCountsAtObservation[i] != 1 {
			t.Errorf("receipt %d = attempt %d, complete %v, close count %d", i, receipt.Attempt, receipt.Complete, closeCountsAtObservation[i])
		}
	}
	if observed[0].StatusCode != http.StatusTooManyRequests || observed[1].StatusCode != http.StatusOK {
		t.Errorf("statuses = %d, %d", observed[0].StatusCode, observed[1].StatusCode)
	}
}

func TestReceiptCopiesAreSafeForConcurrentReaders(t *testing.T) {
	body := newTrackedReadCloser([]byte(`[{"id":"provider-aave","tvl":1,"chainTvls":{"Ethereum":1}}]`), nil)
	client := lifecycleClient(t, http.StatusOK, body)
	receipt, err := client.TVL().GetProtocolsReceipt(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			bodyCopy := receipt.Body()
			if len(bodyCopy) > 0 {
				bodyCopy[0] ^= 0xff
			}
			header := receipt.Header()
			header.Set("X-Test", "mutated")
			_ = receipt.BodySHA256()
		}()
	}
	wg.Wait()
	if got := receipt.Header().Get("X-Test"); got != "" {
		t.Errorf("receipt header was mutated: %q", got)
	}
}

func TestResponseLimitAndVersionMetadata(t *testing.T) {
	if DefaultMaxResponseBodyBytes != 32<<20 {
		t.Errorf("default max response bytes = %d", DefaultMaxResponseBodyBytes)
	}
	if Version != "1.1.0" {
		t.Errorf("Version = %q, want 1.1.0", Version)
	}
	client, err := New()
	if err != nil {
		t.Fatal(err)
	}
	if got := client.cfg.userAgentHeader(); got != "defillama-go/1.1.0" {
		t.Errorf("default User-Agent = %q", got)
	}
	for _, max := range []int64{0, -1, math.MaxInt64} {
		if _, err := New(WithMaxResponseBodyBytes(max)); err == nil {
			t.Errorf("WithMaxResponseBodyBytes(%d) was accepted", max)
		}
	}
}

func lifecycleClient(t *testing.T, status int, body io.ReadCloser, opts ...Option) *Client {
	t.Helper()
	options := append([]Option{
		WithHTTPClient(&http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return &http.Response{StatusCode: status, Header: make(http.Header), Body: body, Request: r}, nil
		})}),
	}, opts...)
	client, err := New(options...)
	if err != nil {
		t.Fatal(err)
	}
	return client
}

func assertResponseBodyTooLarge(t *testing.T, err error, wantLimit int64) {
	t.Helper()
	if !errors.Is(err, ErrResponseBodyTooLarge) {
		t.Fatalf("err = %v, want ErrResponseBodyTooLarge", err)
	}
	var tooLarge *ResponseBodyTooLargeError
	if !errors.As(err, &tooLarge) || tooLarge.Limit != wantLimit {
		t.Fatalf("err = %v, want ResponseBodyTooLargeError limit %d", err, wantLimit)
	}
}

func sha256Hex(body []byte) string {
	digest := sha256.Sum256(body)
	return fmt.Sprintf("%x", digest)
}

type trackedReadCloser struct {
	mu       sync.Mutex
	reader   *bytes.Reader
	closeErr error
	closes   int
}

func newTrackedReadCloser(body []byte, closeErr error) *trackedReadCloser {
	return &trackedReadCloser{reader: bytes.NewReader(body), closeErr: closeErr}
}

func (r *trackedReadCloser) Read(p []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.reader.Read(p)
}

func (r *trackedReadCloser) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closes++
	return r.closeErr
}

func (r *trackedReadCloser) closeCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closes
}

type partialReadCloser struct {
	mu       sync.Mutex
	body     []byte
	readErr  error
	closeErr error
	read     bool
	closes   int
}

func newPartialReadCloser(body []byte, readErr, closeErr error) *partialReadCloser {
	return &partialReadCloser{body: body, readErr: readErr, closeErr: closeErr}
}

func (r *partialReadCloser) Read(p []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.read {
		return 0, io.EOF
	}
	r.read = true
	return copy(p, r.body), r.readErr
}

func (r *partialReadCloser) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.closes++
	return r.closeErr
}

func (r *partialReadCloser) closeCount() int {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.closes
}
