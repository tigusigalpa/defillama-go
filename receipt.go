package defillama

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

// ResponseReceipt is an immutable record of one HTTP response received by the
// SDK. It preserves the response body byte-for-byte and records the selected
// route, its redacted request URL, and transport lifecycle times. It may hold
// an incomplete prefix when the body cannot be fully read or exceeds the
// configured response limit.
//
// Body returns a copy. Decode uses json.Decoder.UseNumber, so numbers decoded
// into interface values remain json.Number rather than float64. Receipts are
// intended for callers that need a lossless boundary around provider-native
// data; callers are responsible for retaining them securely.
type ResponseReceipt struct {
	// Route is the SDK route identifier, for example "GET /protocols".
	Route string
	// SourceURL is the redacted request URL actually selected for this response.
	SourceURL string
	// DocsURL points to the route's official API reference when available.
	DocsURL string
	// StatusCode is the HTTP response status.
	StatusCode int
	// Attempt is the one-based request attempt number for this operation.
	Attempt int
	// CapturedAt is when the SDK received the response, in UTC. It is distinct
	// from any provider timestamp in the response body.
	CapturedAt time.Time
	// CompletedAt is when bounded reading, draining, and closing the original
	// response body completed. It is distinct from provider event time.
	CompletedAt time.Time
	// Complete reports that the original body reached EOF within the configured
	// limit and closed successfully. It does not indicate an HTTP success, valid
	// JSON, or provider-data quality.
	Complete bool

	body   []byte
	header http.Header
}

func newResponseReceipt(routeID, sourceURL, docsURL string, statusCode int, header http.Header, body []byte, attempt int, capturedAt, completedAt time.Time, complete bool) ResponseReceipt {
	return ResponseReceipt{
		Route:       routeID,
		SourceURL:   sourceURL,
		DocsURL:     docsURL,
		StatusCode:  statusCode,
		Attempt:     attempt,
		CapturedAt:  capturedAt,
		CompletedAt: completedAt,
		Complete:    complete,
		body:        append([]byte(nil), body...),
		header:      cloneHeader(header),
	}
}

// Body returns an independent copy of the exact response body bytes.
func (r ResponseReceipt) Body() []byte {
	return append([]byte(nil), r.body...)
}

// Header returns an independent copy of the response headers.
func (r ResponseReceipt) Header() http.Header {
	return cloneHeader(r.header)
}

// BodySHA256 returns the lowercase SHA-256 digest of the exact response body.
func (r ResponseReceipt) BodySHA256() string {
	digest := sha256.Sum256(r.body)
	return hex.EncodeToString(digest[:])
}

// Decode decodes the exact receipt body into out with json.Number preserved.
// It rejects trailing data and multiple JSON values.
func (r ResponseReceipt) Decode(out any) error {
	if out == nil {
		return errors.New("defillama: receipt decode target must not be nil")
	}
	dec := json.NewDecoder(bytes.NewReader(r.body))
	dec.UseNumber()
	if err := dec.Decode(out); err != nil {
		return fmt.Errorf("defillama: decode receipt: %w", err)
	}
	var extra any
	if err := dec.Decode(&extra); err != io.EOF {
		if err == nil {
			err = errors.New("multiple JSON values in receipt")
		}
		return fmt.Errorf("defillama: decode receipt: %w", err)
	}
	return nil
}

func cloneHeader(header http.Header) http.Header {
	if header == nil {
		return nil
	}
	clone := make(http.Header, len(header))
	for key, values := range header {
		clone[key] = append([]string(nil), values...)
	}
	return clone
}

// ReceiptObserver receives one immutable receipt for every HTTP response the
// SDK receives, including retryable responses before a later retry succeeds.
// The observer runs synchronously in the calling request's goroutine.
type ReceiptObserver func(ResponseReceipt)
