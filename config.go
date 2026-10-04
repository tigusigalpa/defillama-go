package defillama

import (
	"crypto/tls"
	"math"
	"net"
	"net/http"
	"time"
)

// DefaultMaxResponseBodyBytes is the 32 MiB ceiling applied to every HTTP
// response body. It is deliberately large enough for protocol and pool lists;
// use WithMaxResponseBodyBytes when an application has a different budget.
const DefaultMaxResponseBodyBytes int64 = 32 << 20

// config holds the immutable client configuration assembled by Option values.
type config struct {
	apiKey           string
	httpClient       *http.Client
	timeout          time.Duration
	retry            RetryPolicy
	userAgent        string
	preferProForFree bool
	baseURLs         map[string]string
	receiptObserver  ReceiptObserver
	maxResponseBytes int64
}

func defaultConfig() config {
	return config{
		timeout:          30 * time.Second,
		retry:            RetryPolicy{MaxAttempts: 1, BaseDelay: 500 * time.Millisecond, MaxDelay: 8 * time.Second, Jitter: true},
		userAgent:        "defillama-go/" + Version,
		maxResponseBytes: DefaultMaxResponseBodyBytes,
	}
}

func (c *config) validate() error {
	if c.timeout <= 0 {
		return &ConfigError{Msg: "timeout must be > 0"}
	}
	if c.retry.MaxAttempts < 1 {
		return &ConfigError{Msg: "retry MaxAttempts must be >= 1"}
	}
	if c.retry.BaseDelay < 0 || c.retry.MaxDelay < 0 {
		return &ConfigError{Msg: "retry delays must be >= 0"}
	}
	if c.maxResponseBytes <= 0 || c.maxResponseBytes == math.MaxInt64 {
		return &ConfigError{Msg: "max response body bytes must be > 0 and below MaxInt64"}
	}
	return nil
}

// userAgentHeader returns the User-Agent sent with every request.
func (c *config) userAgentHeader() string { return c.userAgent }

// Option configures a Client.
type Option func(*config) error

// WithAPIKey sets the DefiLlama Pro API key. The key is embedded as a single
// URL path segment (https://pro-api.llama.fi/{API_KEY}/...) — never a header
// or query parameter — and is redacted from all errors and diagnostics.
// The application decides how to source the secret (e.g. os.Getenv).
func WithAPIKey(key string) Option {
	return func(c *config) error {
		c.apiKey = key
		return nil
	}
}

// WithHTTPClient sets a custom *http.Client for all requests. When unset, a
// default client is built using WithTimeout.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *config) error {
		if hc == nil {
			return &ConfigError{Msg: "http client must not be nil"}
		}
		c.httpClient = hc
		return nil
	}
}

// WithTimeout sets the default request timeout. It is applied to the default
// http.Client; a custom WithHTTPClient manages its own timeouts. Per-request
// deadlines can always be tightened via context.
func WithTimeout(d time.Duration) Option {
	return func(c *config) error {
		if d <= 0 {
			return &ConfigError{Msg: "timeout must be > 0"}
		}
		c.timeout = d
		return nil
	}
}

// WithRetryPolicy configures bounded retries for idempotent GET requests.
// Only transient network errors, HTTP 429 and HTTP 5xx are retried;
// Retry-After is honored. Default: no retries.
func WithRetryPolicy(p RetryPolicy) Option {
	return func(c *config) error {
		if p.MaxAttempts < 1 {
			return &ConfigError{Msg: "retry MaxAttempts must be >= 1"}
		}
		c.retry = p
		return nil
	}
}

// WithUserAgent overrides the default "defillama-go/<version>" User-Agent.
func WithUserAgent(ua string) Option {
	return func(c *config) error {
		if ua == "" {
			return &ConfigError{Msg: "user agent must not be empty"}
		}
		c.userAgent = ua
		return nil
	}
}

// WithPreferProForFree routes the 31 Free endpoints through their official Pro
// equivalents when an API key is configured (/protocols -> /api/protocols,
// prices -> /coins/..., stablecoins -> /stablecoins/..., yields -> /yields/...).
// Default: false — Free routes use their own per-operation origins.
func WithPreferProForFree(v bool) Option {
	return func(c *config) error {
		c.preferProForFree = v
		return nil
	}
}

// WithBaseURLsForTesting overrides default base URLs. For tests and advanced
// use only: keys are the default origins ("https://api.llama.fi",
// "https://coins.llama.fi", "https://stablecoins.llama.fi",
// "https://yields.llama.fi", "https://pro-api.llama.fi").
func WithBaseURLsForTesting(m map[string]string) Option {
	return func(c *config) error {
		// Keep a private copy so changes made by the caller after New cannot
		// alter a live client or introduce a data race.
		c.baseURLs = make(map[string]string, len(m))
		for base, override := range m {
			c.baseURLs[base] = override
		}
		return nil
	}
}

// WithReceiptObserver receives an immutable ResponseReceipt for every HTTP
// response the SDK receives, including retryable responses before a later
// attempt succeeds. Use it when your application needs to retain provider
// receipts outside the SDK's ordinary typed response boundary.
//
// The observer runs synchronously in the request goroutine. It receives a
// redacted URL and defensive copies from Body and Header; avoid logging or
// persisting receipt bodies unless your application's data policy allows it.
func WithReceiptObserver(observer ReceiptObserver) Option {
	return func(c *config) error {
		if observer == nil {
			return &ConfigError{Msg: "receipt observer must not be nil"}
		}
		c.receiptObserver = observer
		return nil
	}
}

// WithMaxResponseBodyBytes sets the maximum number of response bytes retained
// and decoded for each request. The default is DefaultMaxResponseBodyBytes
// (32 MiB). An over-limit response returns a ResponseBodyTooLargeError and is
// never retried.
func WithMaxResponseBodyBytes(limit int64) Option {
	return func(c *config) error {
		if limit <= 0 || limit == math.MaxInt64 {
			return &ConfigError{Msg: "max response body bytes must be > 0 and below MaxInt64"}
		}
		c.maxResponseBytes = limit
		return nil
	}
}

// httpClient returns the configured client or builds a sane default.
func (c *config) http() *http.Client {
	if c.httpClient != nil {
		return c.httpClient
	}
	return &http.Client{
		Timeout: c.timeout,
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
			TLSClientConfig:     &tls.Config{MinVersion: tls.VersionTLS12},
			MaxIdleConns:        100,
			IdleConnTimeout:     90 * time.Second,
		},
	}
}

// RetryPolicy configures bounded retries for GET requests.
type RetryPolicy struct {
	// MaxAttempts is the total number of attempts including the first (1 = off).
	MaxAttempts int
	// BaseDelay is the delay before the first retry; it doubles per attempt.
	BaseDelay time.Duration
	// MaxDelay caps a single delay.
	MaxDelay time.Duration
	// Jitter adds +/-50% randomization to each computed delay.
	Jitter bool
}

// delay computes the backoff before retry number n (1-based).
func (p RetryPolicy) delay(n int) time.Duration {
	d := p.BaseDelay << (n - 1)
	if d <= 0 || d > p.MaxDelay {
		d = p.MaxDelay
	}
	if p.Jitter {
		half := d / 2
		d = half + time.Duration(rand63n(int64(half)+1))
	}
	return d
}
