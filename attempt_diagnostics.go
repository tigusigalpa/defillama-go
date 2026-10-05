package defillama

import (
	"context"
	"net/url"
	"time"
)

// AttemptOutcome describes whether one observed HTTP attempt completed
// successfully. It is independent from response-body completeness and from the
// outcome of later attempts in the same operation.
type AttemptOutcome string

const (
	// AttemptSucceeded means the attempt completed without an SDK error.
	AttemptSucceeded AttemptOutcome = "succeeded"
	// AttemptFailed means the attempt completed with an SDK error. Err exposes
	// the original typed or wrapped cause.
	AttemptFailed AttemptOutcome = "failed"
)

// AttemptDiagnostic is immutable evidence for one observed SDK attempt. It
// never contains request headers, response bodies, or the raw request URL.
// TargetURL has user-info removed and query values redacted.
//
// Err is nil for a successful attempt. For a failed attempt it preserves the
// original cause, so errors.Is and errors.As continue to work. Diagnostics do
// not format, log, or serialize Err automatically.
type AttemptDiagnostic struct {
	// OperationID identifies one SDK operation within a Client. It lets a shared
	// observer separate concurrent operations for the same route and target.
	OperationID uint64
	// Route is the registered SDK route identifier, for example "GET /protocols".
	Route string
	// TargetURL is a redacted diagnostic target, not a replayable request URL.
	TargetURL string
	// DocsURL points to the route's official API reference when available.
	DocsURL string
	// Attempt is the one-based attempt number within OperationID.
	Attempt int
	// StartedAt is the local UTC time immediately before this attempt is built.
	StartedAt time.Time
	// CompletedAt is the local UTC time after the attempt and, when applicable,
	// its original response-body lifecycle have completed.
	CompletedAt time.Time
	// ResponseReceived reports whether the attempt received an HTTP response.
	ResponseReceived bool
	// StatusCode is the observed HTTP status when ResponseReceived is true.
	// It is zero when no response was received.
	StatusCode int
	// ResponseComplete reports whether an observed response body reached EOF
	// within the configured limit and closed successfully.
	ResponseComplete bool
	// Outcome describes this attempt only; it does not predict a later retry or
	// the final operation result.
	Outcome AttemptOutcome
	// Err is the real SDK error for this attempt, when known.
	Err error
}

// AttemptObserver receives diagnostics for one operation context. It runs
// synchronously in the calling goroutine after any received response body has
// been finalized. It is never called for work that did not reach an SDK
// attempt, such as route validation failures.
//
// Use a fresh observer context for each operation. The SDK retains no attempt
// history, so diagnostics stay bounded by the caller's callback behavior.
type AttemptObserver func(AttemptDiagnostic)

type attemptObserverContextKey struct{}

// WithAttemptObserver returns a child context that receives one
// AttemptDiagnostic for every SDK attempt performed with it. This includes
// response-less transport failures, which remain distinct from HTTP receipts.
// Passing nil leaves ctx unchanged.
func WithAttemptObserver(ctx context.Context, observer AttemptObserver) context.Context {
	if observer == nil {
		return ctx
	}
	return context.WithValue(ctx, attemptObserverContextKey{}, observer)
}

func attemptObserverFromContext(ctx context.Context) AttemptObserver {
	observer, _ := ctx.Value(attemptObserverContextKey{}).(AttemptObserver)
	return observer
}

func newAttemptDiagnostic(operationID uint64, r route, targetURL string, attempt int, startedAt, completedAt time.Time, responseReceived bool, statusCode int, responseComplete bool, err error) AttemptDiagnostic {
	outcome := AttemptSucceeded
	if err != nil {
		outcome = AttemptFailed
	}
	return AttemptDiagnostic{
		OperationID:      operationID,
		Route:            r.Method + " " + r.Path,
		TargetURL:        diagnosticTargetURL(targetURL),
		DocsURL:          r.DocsURL,
		Attempt:          attempt,
		StartedAt:        startedAt,
		CompletedAt:      completedAt,
		ResponseReceived: responseReceived,
		StatusCode:       statusCode,
		ResponseComplete: responseComplete,
		Outcome:          outcome,
		Err:              err,
	}
}

func diagnosticTargetURL(target string) string {
	u, err := url.Parse(target)
	if err != nil {
		return ""
	}
	u.User = nil
	query := u.Query()
	for key, values := range query {
		for i := range values {
			values[i] = "{REDACTED}"
		}
		query[key] = values
	}
	u.RawQuery = query.Encode()
	return u.String()
}
