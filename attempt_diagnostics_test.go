package defillama

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"
)

func TestAttemptObserverRecordsRetryOutcomes(t *testing.T) {
	for _, statusCode := range []int{http.StatusTooManyRequests, http.StatusServiceUnavailable} {
		t.Run(http.StatusText(statusCode), func(t *testing.T) {
			calls := 0
			client, err := New(
				WithHTTPClient(&http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
					calls++
					if calls == 1 {
						return diagnosticResponse(r, statusCode, `{"error":"retry"}`), nil
					}
					return diagnosticResponse(r, http.StatusOK, `[]`), nil
				})}),
				WithRetryPolicy(RetryPolicy{MaxAttempts: 2, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}),
			)
			if err != nil {
				t.Fatal(err)
			}

			ctx, diagnostics := captureAttemptDiagnostics(context.Background())
			if _, err := client.TVL().GetProtocols(ctx); err != nil {
				t.Fatal(err)
			}
			got := diagnostics()
			if len(got) != 2 {
				t.Fatalf("diagnostics = %d, want 2", len(got))
			}
			assertAttempt(t, got[0], 1, statusCode, true, true, AttemptFailed)
			assertAttempt(t, got[1], 2, http.StatusOK, true, true, AttemptSucceeded)
			if got[0].OperationID == 0 || got[0].OperationID != got[1].OperationID {
				t.Errorf("operation IDs = %d, %d", got[0].OperationID, got[1].OperationID)
			}
			if got[0].CompletedAt.Before(got[0].StartedAt) || got[1].CompletedAt.Before(got[1].StartedAt) {
				t.Error("attempt completion precedes its start")
			}
			if got[0].Err == nil || got[1].Err != nil {
				t.Errorf("attempt errors = %v, %v", got[0].Err, got[1].Err)
			}
		})
	}
}

func TestAttemptObserverRetainsIndependentExhaustedRetryErrors(t *testing.T) {
	client, err := New(
		WithHTTPClient(&http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			return diagnosticResponse(r, http.StatusBadGateway, `{"error":"upstream"}`), nil
		})}),
		WithRetryPolicy(RetryPolicy{MaxAttempts: 2, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, diagnostics := captureAttemptDiagnostics(context.Background())
	_, err = client.TVL().GetProtocols(ctx)
	var terminal *APIError
	if !errors.As(err, &terminal) || terminal.StatusCode != http.StatusBadGateway {
		t.Fatalf("terminal error = %v, want HTTP 502 APIError", err)
	}
	got := diagnostics()
	if len(got) != 2 {
		t.Fatalf("diagnostics = %d, want 2", len(got))
	}
	for index, diagnostic := range got {
		assertAttempt(t, diagnostic, index+1, http.StatusBadGateway, true, true, AttemptFailed)
		var apiErr *APIError
		if !errors.As(diagnostic.Err, &apiErr) || apiErr.StatusCode != http.StatusBadGateway {
			t.Errorf("attempt %d error = %v, want HTTP 502 APIError", index+1, diagnostic.Err)
		}
	}
	if got[0].Err == got[1].Err {
		t.Error("attempt errors share the same instance")
	}
}

func TestAttemptObserverSeparatesResponseLessTerminalFailure(t *testing.T) {
	calls := 0
	var receiptCalls int
	client, err := New(
		WithHTTPClient(&http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
			calls++
			if calls == 1 {
				return diagnosticResponse(r, http.StatusTooManyRequests, `{"error":"slow down"}`), nil
			}
			return nil, io.ErrUnexpectedEOF
		})}),
		WithRetryPolicy(RetryPolicy{MaxAttempts: 2, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}),
		WithReceiptObserver(func(ResponseReceipt) { receiptCalls++ }),
	)
	if err != nil {
		t.Fatal(err)
	}

	ctx, diagnostics := captureAttemptDiagnostics(context.Background())
	receipt, err := client.TVL().GetProtocolsReceipt(ctx)
	if !errors.Is(err, io.ErrUnexpectedEOF) {
		t.Fatalf("err = %v, want io.ErrUnexpectedEOF", err)
	}
	if receipt != nil {
		t.Errorf("receipt = %#v, want nil for response-less terminal failure", receipt)
	}
	if receiptCalls != 1 {
		t.Errorf("receipt observer calls = %d, want 1", receiptCalls)
	}
	got := diagnostics()
	if len(got) != 2 {
		t.Fatalf("diagnostics = %d, want 2", len(got))
	}
	assertAttempt(t, got[0], 1, http.StatusTooManyRequests, true, true, AttemptFailed)
	assertAttempt(t, got[1], 2, 0, false, false, AttemptFailed)
	if !errors.Is(got[1].Err, io.ErrUnexpectedEOF) {
		t.Errorf("terminal attempt error = %v, want io.ErrUnexpectedEOF", got[1].Err)
	}
}

func TestAttemptObserverPreservesCancellationCauses(t *testing.T) {
	t.Run("response-less cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		client, err := New(WithHTTPClient(&http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
			cancel()
			return nil, context.Canceled
		})}))
		if err != nil {
			t.Fatal(err)
		}
		ctx, diagnostics := captureAttemptDiagnostics(ctx)
		_, err = client.TVL().GetProtocols(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("err = %v, want context.Canceled", err)
		}
		got := diagnostics()
		if len(got) != 1 {
			t.Fatalf("diagnostics = %d, want 1", len(got))
		}
		assertAttempt(t, got[0], 1, 0, false, false, AttemptFailed)
		if !errors.Is(got[0].Err, context.Canceled) {
			t.Errorf("diagnostic error = %v, want context.Canceled", got[0].Err)
		}
	})

	t.Run("response-less deadline", func(t *testing.T) {
		client, err := New(WithHTTPClient(&http.Client{Transport: roundTripFunc(func(request *http.Request) (*http.Response, error) {
			<-request.Context().Done()
			return nil, request.Context().Err()
		})}))
		if err != nil {
			t.Fatal(err)
		}
		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
		defer cancel()
		ctx, diagnostics := captureAttemptDiagnostics(ctx)
		_, err = client.TVL().GetProtocols(ctx)
		if !errors.Is(err, context.DeadlineExceeded) {
			t.Fatalf("err = %v, want context.DeadlineExceeded", err)
		}
		got := diagnostics()
		if len(got) != 1 {
			t.Fatalf("diagnostics = %d, want 1", len(got))
		}
		assertAttempt(t, got[0], 1, 0, false, false, AttemptFailed)
		if !errors.Is(got[0].Err, context.DeadlineExceeded) {
			t.Errorf("diagnostic error = %v, want context.DeadlineExceeded", got[0].Err)
		}
	})

	t.Run("cancellation while reading body", func(t *testing.T) {
		body := newPartialReadCloser([]byte(`[`), context.Canceled, nil)
		client := lifecycleClient(t, http.StatusOK, body)
		ctx, diagnostics := captureAttemptDiagnostics(context.Background())
		_, err := client.TVL().GetProtocols(ctx)
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("err = %v, want context.Canceled", err)
		}
		got := diagnostics()
		if len(got) != 1 {
			t.Fatalf("diagnostics = %d, want 1", len(got))
		}
		assertAttempt(t, got[0], 1, http.StatusOK, true, false, AttemptFailed)
		if !errors.Is(got[0].Err, context.Canceled) {
			t.Errorf("diagnostic error = %v, want context.Canceled", got[0].Err)
		}
	})
}

func TestAttemptObserverIsolatesConcurrentOperations(t *testing.T) {
	client, err := New(WithHTTPClient(&http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return diagnosticResponse(r, http.StatusOK, `[]`), nil
	})}))
	if err != nil {
		t.Fatal(err)
	}

	const operations = 16
	operationIDs := make(chan uint64, operations)
	var wg sync.WaitGroup
	for range operations {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ctx, diagnostics := captureAttemptDiagnostics(context.Background())
			if _, err := client.TVL().GetProtocols(ctx); err != nil {
				t.Error(err)
				return
			}
			got := diagnostics()
			if len(got) != 1 {
				t.Errorf("diagnostics = %d, want 1", len(got))
				return
			}
			operationIDs <- got[0].OperationID
		}()
	}
	wg.Wait()
	close(operationIDs)
	seen := make(map[uint64]struct{}, operations)
	for operationID := range operationIDs {
		if operationID == 0 {
			t.Error("operation ID must not be zero")
		}
		if _, exists := seen[operationID]; exists {
			t.Errorf("operation ID %d mixed concurrent diagnostics", operationID)
		}
		seen[operationID] = struct{}{}
	}
	if len(seen) != operations {
		t.Errorf("operation IDs = %d, want %d", len(seen), operations)
	}
}

func TestAttemptDiagnosticTargetRedaction(t *testing.T) {
	target := diagnosticTargetURL("https://user:password@example.test/path?api_key=secret&token=value&safe=visible")
	if strings.Contains(target, "user") || strings.Contains(target, "password") || strings.Contains(target, "secret") || strings.Contains(target, "value") {
		t.Errorf("diagnostic target leaked a secret: %q", target)
	}
	if strings.Contains(target, "visible") || !strings.Contains(target, "safe=%7BREDACTED%7D") || !strings.Contains(target, "api_key=%7BREDACTED%7D") {
		t.Errorf("diagnostic target = %q", target)
	}
}

func captureAttemptDiagnostics(ctx context.Context) (context.Context, func() []AttemptDiagnostic) {
	var mu sync.Mutex
	var captured []AttemptDiagnostic
	ctx = WithAttemptObserver(ctx, func(diagnostic AttemptDiagnostic) {
		mu.Lock()
		defer mu.Unlock()
		captured = append(captured, diagnostic)
	})
	return ctx, func() []AttemptDiagnostic {
		mu.Lock()
		defer mu.Unlock()
		return append([]AttemptDiagnostic(nil), captured...)
	}
}

func diagnosticResponse(request *http.Request, statusCode int, body string) *http.Response {
	return &http.Response{
		StatusCode: statusCode,
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader(body)),
		Request:    request,
	}
}

func assertAttempt(t *testing.T, diagnostic AttemptDiagnostic, attempt, statusCode int, responseReceived, responseComplete bool, outcome AttemptOutcome) {
	t.Helper()
	if diagnostic.Route != "GET /protocols" {
		t.Errorf("route = %q", diagnostic.Route)
	}
	if diagnostic.TargetURL == "" || strings.Contains(diagnostic.TargetURL, "SECRET_TEST_KEY_1") {
		t.Errorf("diagnostic target = %q", diagnostic.TargetURL)
	}
	if diagnostic.Attempt != attempt || diagnostic.StatusCode != statusCode || diagnostic.ResponseReceived != responseReceived || diagnostic.ResponseComplete != responseComplete || diagnostic.Outcome != outcome {
		t.Errorf("attempt = %+v", diagnostic)
	}
}
