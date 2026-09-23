package defillama

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

// ConfigError reports invalid client configuration. Returned by New before any
// network I/O takes place.
type ConfigError struct {
	Msg string
}

// Error implements error.
func (e *ConfigError) Error() string { return "defillama: config: " + e.Msg }

// ProAPIKeyRequiredError is returned locally — before any HTTP request is
// built — when a Pro endpoint is called without an API key.
type ProAPIKeyRequiredError struct {
	// Route is the registry route id ("GET /path").
	Route string
}

// Error implements error.
func (e *ProAPIKeyRequiredError) Error() string {
	return fmt.Sprintf("defillama: endpoint %q requires a Pro API key (use WithAPIKey)", e.Route)
}

// TransportError wraps a network-level failure (DNS, TLS, timeout, ...).
// Unwrap exposes the underlying error; errors.Is/As work through it.
type TransportError struct {
	// URL is the redacted request URL (API key replaced by {API_KEY}).
	URL string
	Err error
}

// Error implements error.
func (e *TransportError) Error() string { return "defillama: " + e.URL + ": " + e.Err.Error() }

// Unwrap returns the wrapped network error.
func (e *TransportError) Unwrap() error { return e.Err }

// DecodeError reports a response body that is not valid JSON.
type DecodeError struct {
	// URL is the redacted request URL.
	URL string
	Err error
}

// Error implements error.
func (e *DecodeError) Error() string {
	return "defillama: cannot decode JSON response from " + e.URL + ": " + e.Err.Error()
}

// Unwrap returns the JSON decoding error.
func (e *DecodeError) Unwrap() error { return e.Err }

// apiErrorBodyLimit bounds the response body stored on APIError.
const apiErrorBodyLimit = 4096

// APIError is a non-2xx HTTP response. URL is always the redacted form.
type APIError struct {
	StatusCode int
	// URL is the redacted request URL (API key replaced by {API_KEY}).
	URL string
	// Header is a copy of the response headers with API key values redacted.
	Header http.Header
	// Body is the response body truncated to 4096 bytes with API key values
	// redacted. It is safe to include in application diagnostics.
	Body []byte
}

// Error implements error.
func (e *APIError) Error() string {
	return fmt.Sprintf("defillama: HTTP %d for %s", e.StatusCode, e.URL)
}

// NotFoundError is an HTTP 404 response.
type NotFoundError struct {
	*APIError
}

// Error implements error.
func (e *NotFoundError) Error() string { return e.APIError.Error() }

// RateLimitError is an HTTP 429 response. RetryAfter is the parsed Retry-After
// value, or zero when the header was absent/unparseable.
type RateLimitError struct {
	*APIError
	RetryAfter time.Duration
}

// Error implements error.
func (e *RateLimitError) Error() string {
	if e.RetryAfter > 0 {
		return fmt.Sprintf("%s (retry after %s)", e.APIError.Error(), e.RetryAfter)
	}
	return e.APIError.Error()
}

// redactHeader returns a copy of h with the API key removed from every value.
func redactHeader(h http.Header, apiKey string) http.Header {
	redacted := h.Clone()
	for name, values := range redacted {
		for i, value := range values {
			values[i] = redactString(value, apiKey)
		}
		redacted[name] = values
	}
	return redacted
}

// redactBytes removes both the literal and path-escaped API key from an HTTP
// error body. Error payloads are diagnostic data and must be safe to log.
func redactBytes(body []byte, apiKey string) []byte {
	if apiKey == "" || len(body) == 0 {
		return body
	}
	redacted := append([]byte(nil), body...)
	for _, secret := range apiKeyForms(apiKey) {
		redacted = bytes.ReplaceAll(redacted, []byte(secret), []byte("{API_KEY}"))
	}
	return redacted
}

func redactString(value, apiKey string) string {
	if apiKey == "" || value == "" {
		return value
	}
	for _, secret := range apiKeyForms(apiKey) {
		value = string(bytes.ReplaceAll([]byte(value), []byte(secret), []byte("{API_KEY}")))
	}
	return value
}

func apiKeyForms(apiKey string) []string {
	escaped := url.PathEscape(apiKey)
	if escaped == apiKey {
		return []string{apiKey}
	}
	return []string{apiKey, escaped}
}

// redactedTransportCause keeps a transport cause discoverable through
// errors.Is/As while making its diagnostic message safe to log.
type redactedTransportCause struct {
	err     error
	message string
}

func (e *redactedTransportCause) Error() string { return e.message }

func (e *redactedTransportCause) Unwrap() error { return e.err }
