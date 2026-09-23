package defillama

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// testServer returns an httptest server plus a client wired to it for all
// origins, and a request log.
func testServer(t *testing.T, handler http.HandlerFunc, opts ...Option) (*httptest.Server, *Client, *requestLog) {
	t.Helper()
	log := &requestLog{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.add(r)
		handler(w, r)
	}))
	t.Cleanup(srv.Close)
	opts = append([]Option{
		WithAPIKey("SECRET_TEST_KEY_1"),
		WithBaseURLsForTesting(map[string]string{
			"https://api.llama.fi":         srv.URL,
			"https://coins.llama.fi":       srv.URL,
			"https://stablecoins.llama.fi": srv.URL,
			"https://yields.llama.fi":      srv.URL,
			proBase:                        srv.URL,
		}),
	}, opts...)
	c, err := New(opts...)
	if err != nil {
		t.Fatal(err)
	}
	return srv, c, log
}

type requestLog struct {
	mu   sync.Mutex
	reqs []*http.Request
}

func (l *requestLog) add(r *http.Request) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.reqs = append(l.reqs, r)
}

func (l *requestLog) all() []*http.Request {
	l.mu.Lock()
	defer l.mu.Unlock()
	return append([]*http.Request{}, l.reqs...)
}

// okJSON responds with JSON null, which decodes cleanly into any target type.
func okJSON(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, `null`)
}

func TestHeadersSent(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	if _, err := c.TVL().GetProtocols(context.Background()); err != nil {
		t.Fatal(err)
	}
	r := log.all()[0]
	if got := r.Header.Get("Accept"); got != "application/json" {
		t.Errorf("Accept = %q", got)
	}
	if got := r.Header.Get("User-Agent"); !strings.HasPrefix(got, "defillama-go/") {
		t.Errorf("User-Agent = %q", got)
	}
}

func TestFreeRouteUsesPath(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	if _, err := c.TVL().GetProtocol(context.Background(), "aave"); err != nil {
		t.Fatal(err)
	}
	if got := log.all()[0].URL.Path; got != "/protocol/aave" {
		t.Errorf("path = %q", got)
	}
}

func TestProRouteKeyInPath(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	if _, err := c.Emissions().GetEmissions(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got := log.all()[0].URL.Path; got != "/SECRET_TEST_KEY_1/api/emissions" {
		t.Errorf("path = %q", got)
	}
}

func TestPreferProForFreeOverride(t *testing.T) {
	_, c, log := testServer(t, okJSON, WithPreferProForFree(true))
	if _, err := c.TVL().GetProtocols(context.Background()); err != nil {
		t.Fatal(err)
	}
	if got := log.all()[0].URL.Path; got != "/SECRET_TEST_KEY_1/api/protocols" {
		t.Errorf("path = %q", got)
	}
}

func TestProRouteWithoutKeyFailsLocally(t *testing.T) {
	srv, _, log := testServer(t, okJSON)
	c2, err := New(WithBaseURLsForTesting(map[string]string{
		"https://api.llama.fi": srv.URL,
		proBase:                srv.URL,
	}))
	if err != nil {
		t.Fatal(err)
	}
	var target *ProAPIKeyRequiredError
	if _, err := c2.Emissions().GetEmissions(context.Background()); !errors.As(err, &target) {
		t.Fatalf("err = %v, want ProAPIKeyRequiredError", err)
	}
	if len(log.all()) != 0 {
		t.Error("request was sent despite missing key")
	}
}

func TestCoinsEscaping(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	if _, err := c.Prices().GetCurrentPrices(context.Background(), []string{
		"coingecko:bitcoin", "ethereum:0xabc",
	}); err != nil {
		t.Fatal(err)
	}
	got := log.all()[0].URL.EscapedPath()
	wantA := "/prices/current/coingecko:bitcoin,ethereum:0xabc"
	wantB := "/prices/current/coingecko%3Abitcoin,ethereum%3A0xabc"
	if got != wantA && got != wantB {
		t.Errorf("escaped path = %q, want %q or %q", got, wantA, wantB)
	}
}

func TestEmptyCoinListFailsBeforeRequest(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	_, err := c.Prices().GetCurrentPrices(context.Background(), nil)
	if err == nil || !strings.Contains(err.Error(), "path parameter \"coins\" must not be empty") {
		t.Fatalf("err = %v, want an empty coins error", err)
	}
	if len(log.all()) != 0 {
		t.Error("request was sent with an empty coins path parameter")
	}
}

func TestBoolQueryLowercase(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	_, err := c.Volumes().GetDEXOverview(context.Background(), OverviewOptions{
		ExcludeTotalDataChart:          true,
		ExcludeTotalDataChartBreakdown: false,
	})
	if err != nil {
		t.Fatal(err)
	}
	q := log.all()[0].URL.Query()
	if q.Get("excludeTotalDataChart") != "true" || q.Get("excludeTotalDataChartBreakdown") != "false" {
		t.Errorf("query = %s", log.all()[0].URL.RawQuery)
	}
}

func TestBatchHistoricalJSONQuery(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	_, err := c.Prices().GetBatchHistoricalPrices(context.Background(), map[string][]int64{
		"coingecko:bitcoin": {1600000000, 1600000100},
	})
	if err != nil {
		t.Fatal(err)
	}
	got := log.all()[0].URL.Query().Get("coins")
	if got != `{"coingecko:bitcoin":[1600000000,1600000100]}` {
		t.Errorf("coins = %q", got)
	}
}

func TestRequiredQueryEnforced(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	_, err := c.Volumes().GetDEXOverview(context.Background(), OverviewOptions{})
	if err != nil {
		t.Fatal(err)
	}
	q := log.all()[0].URL.Query()
	if !q.Has("excludeTotalDataChart") || !q.Has("excludeTotalDataChartBreakdown") {
		t.Errorf("required query flags missing: %s", log.all()[0].URL.RawQuery)
	}
}

func TestRequiredStringQueryRejectsEmptyValues(t *testing.T) {
	_, c, log := testServer(t, okJSON)

	_, err := c.PreIPO().GetValuations(context.Background(), PreIPOQuery{})
	if err == nil || !strings.Contains(err.Error(), `required query parameter "company"`) {
		t.Errorf("empty company: err = %v", err)
	}
	_, err = c.Equities().GetOHLCV(context.Background(), EquitiesQuery{Ticker: "NVDA"})
	if err == nil || !strings.Contains(err.Error(), `required query parameter "country"`) {
		t.Errorf("empty country: err = %v", err)
	}
	if got := len(log.all()); got != 0 {
		t.Errorf("sent %d request(s) with empty required query parameters", got)
	}
}

func TestEnumValidation(t *testing.T) {
	_, c, log := testServer(t, okJSON)
	opts := HistoryRangeOptions{Range: strPtr("bogus")}
	if _, err := c.Yields().GetEarnPoolHistory(context.Background(), "pool1", &opts); err == nil {
		t.Fatal("expected enum validation error")
	}
	if _, err := c.TVL().GetProtocolTVLChart(context.Background(), "aave",
		&TVLChartOptions{Key: strPtr("bogus")}); err == nil {
		t.Fatal("expected enum validation error for key")
	}
	if _, err := c.Dimensions().GetChart(context.Background(), "bogus", nil); err == nil {
		t.Fatal("expected enum validation error for metric path parameter")
	}
	if got := len(log.all()); got != 0 {
		t.Errorf("sent %d request(s) with invalid enum values", got)
	}
}

func strPtr(s string) *string { return &s }

func TestNotFound(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, `{"error":"nope"}`)
	})
	var nf *NotFoundError
	_, err := c.TVL().GetProtocol(context.Background(), "zzz")
	if !errors.As(err, &nf) {
		t.Fatalf("err = %v (%T), want NotFoundError", err, err)
	}
	if nf.StatusCode != 404 {
		t.Errorf("status = %d", nf.StatusCode)
	}
}

func TestRateLimitRetryAfter(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Retry-After", "42")
		w.WriteHeader(http.StatusTooManyRequests)
	})
	var rl *RateLimitError
	_, err := c.TVL().GetProtocols(context.Background())
	if !errors.As(err, &rl) {
		t.Fatalf("err = %v (%T)", err, err)
	}
	if rl.RetryAfter != 42*time.Second {
		t.Errorf("RetryAfter = %v", rl.RetryAfter)
	}
}

func TestServerError(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, `bad gateway`)
	})
	var ae *APIError
	_, err := c.TVL().GetProtocols(context.Background())
	if !errors.As(err, &ae) {
		t.Fatalf("err = %v (%T)", err, err)
	}
	if ae.StatusCode != 502 || string(ae.Body) != "bad gateway" {
		t.Errorf("api error = %+v", ae)
	}
}

func TestMalformedJSON(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `not json{`)
	})
	var de *DecodeError
	_, err := c.TVL().GetProtocols(context.Background())
	if !errors.As(err, &de) {
		t.Fatalf("err = %v (%T)", err, err)
	}
}

func TestTrailingGarbageIsDecodeError(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `[] not-json`)
	})
	_, err := c.TVL().GetProtocols(context.Background())
	var decodeErr *DecodeError
	if !errors.As(err, &decodeErr) {
		t.Fatalf("err = %v (%T), want DecodeError", err, err)
	}
}

func TestAPIKeyRedaction(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})
	_, err := c.Emissions().GetEmissions(context.Background())
	if err == nil {
		t.Fatal("expected error")
	}
	msg := fmt.Sprintf("%+v", err)
	if strings.Contains(msg, "SECRET_TEST_KEY_1") {
		t.Errorf("api key leaked in error: %s", msg)
	}
	if !strings.Contains(msg, "{API_KEY}") {
		t.Errorf("redacted placeholder missing: %s", msg)
	}
	var ae *APIError
	if errors.As(err, &ae) && strings.Contains(ae.URL, "SECRET_TEST_KEY_1") {
		t.Errorf("api key in error URL: %s", ae.URL)
	}
}

func TestAPIErrorRedactsResponseDiagnostics(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://example.test/SECRET_TEST_KEY_1/next")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"detail":"SECRET_TEST_KEY_1"}`)
	})
	_, err := c.Emissions().GetEmissions(context.Background())
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err = %v (%T), want APIError", err, err)
	}
	if got := apiErr.Header.Get("Location"); strings.Contains(got, "SECRET_TEST_KEY_1") || !strings.Contains(got, "{API_KEY}") {
		t.Errorf("Location = %q, want a redacted value", got)
	}
	if got := string(apiErr.Body); strings.Contains(got, "SECRET_TEST_KEY_1") || !strings.Contains(got, "{API_KEY}") {
		t.Errorf("Body = %q, want a redacted value", got)
	}
}

func TestTransportErrorRedactsWrappedURL(t *testing.T) {
	client := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
		return nil, &url.Error{Op: "Get", URL: r.URL.String(), Err: fmt.Errorf("dial failed for %s", r.URL)}
	})}
	c, err := New(WithAPIKey("SECRET_TEST_KEY_1"), WithHTTPClient(client))
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.Emissions().GetEmissions(context.Background())
	var transportErr *TransportError
	if !errors.As(err, &transportErr) {
		t.Fatalf("err = %v (%T), want TransportError", err, err)
	}
	var urlErr *url.Error
	if !errors.As(err, &urlErr) {
		t.Fatalf("err = %v, want wrapped url.Error", err)
	}
	if strings.Contains(urlErr.URL, "SECRET_TEST_KEY_1") || !strings.Contains(urlErr.URL, "{API_KEY}") {
		t.Errorf("wrapped URL = %q, want a redacted URL", urlErr.URL)
	}
	if got := fmt.Sprintf("%+v", err); strings.Contains(got, "SECRET_TEST_KEY_1") {
		t.Errorf("api key leaked in error: %s", got)
	}
	if got := transportErr.Err.Error(); strings.Contains(got, "SECRET_TEST_KEY_1") {
		t.Errorf("api key leaked in TransportError.Err: %s", got)
	}
}

func TestInvalidProBaseURLDoesNotLeakKey(t *testing.T) {
	c, err := New(
		WithAPIKey("SECRET_TEST_KEY_1"),
		WithBaseURLsForTesting(map[string]string{proBase: "http://%invalid"}),
	)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Emissions().GetEmissions(context.Background())
	var transportErr *TransportError
	if !errors.As(err, &transportErr) {
		t.Fatalf("err = %v (%T), want TransportError", err, err)
	}
	if got := fmt.Sprintf("%+v", err); strings.Contains(got, "SECRET_TEST_KEY_1") {
		t.Errorf("api key leaked from request construction: %s", got)
	}
}

func TestProRedirectCannotSendKeyToAnotherOrigin(t *testing.T) {
	var redirected atomic.Int32
	other := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirected.Add(1)
		okJSON(w, r)
	}))
	t.Cleanup(other.Close)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", other.URL+"/SECRET_TEST_KEY_1/api/emissions")
		w.WriteHeader(http.StatusFound)
	}))
	t.Cleanup(server.Close)
	c, err := New(
		WithAPIKey("SECRET_TEST_KEY_1"),
		WithBaseURLsForTesting(map[string]string{proBase: server.URL}),
	)
	if err != nil {
		t.Fatal(err)
	}
	_, err = c.Emissions().GetEmissions(context.Background())
	var apiErr *APIError
	if !errors.As(err, &apiErr) || apiErr.StatusCode != http.StatusFound {
		t.Fatalf("err = %v (%T), want HTTP 302 APIError", err, err)
	}
	if got := redirected.Load(); got != 0 {
		t.Errorf("redirect target received %d request(s)", got)
	}
	if got := apiErr.Header.Get("Location"); strings.Contains(got, "SECRET_TEST_KEY_1") {
		t.Errorf("api key leaked in Location header: %s", got)
	}
}

func TestProRedirectOnSameOriginStillWorks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/SECRET_TEST_KEY_1/api/emissions" {
			http.Redirect(w, r, "/SECRET_TEST_KEY_1/api/emissions/latest", http.StatusFound)
			return
		}
		fmt.Fprint(w, `[]`)
	}))
	t.Cleanup(server.Close)
	c, err := New(
		WithAPIKey("SECRET_TEST_KEY_1"),
		WithBaseURLsForTesting(map[string]string{proBase: server.URL}),
	)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := c.Emissions().GetEmissions(context.Background()); err != nil {
		t.Fatalf("same-origin redirect failed: %v", err)
	}
}

func TestFreeErrorBodyIsNotChangedByUnusedProKey(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "bad parameter")
	}, WithAPIKey("a"))
	_, err := c.TVL().GetProtocols(context.Background())
	var apiErr *APIError
	if !errors.As(err, &apiErr) {
		t.Fatalf("err = %v (%T), want APIError", err, err)
	}
	if got := string(apiErr.Body); got != "bad parameter" {
		t.Errorf("Free error body changed: %q", got)
	}
}

func TestRetryOn5xx(t *testing.T) {
	var calls atomic.Int32
	_, c, log := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		if calls.Add(1) == 1 {
			w.WriteHeader(503)
			return
		}
		fmt.Fprint(w, `[]`)
	}, WithRetryPolicy(RetryPolicy{MaxAttempts: 3, BaseDelay: time.Millisecond}))
	if _, err := c.TVL().GetProtocols(context.Background()); err != nil {
		t.Fatal(err)
	}
	if len(log.all()) != 2 {
		t.Errorf("requests = %d, want 2", len(log.all()))
	}
}

func TestRetryOnlyTransientTransportErrors(t *testing.T) {
	for _, tt := range []struct {
		name      string
		failure   error
		wantCalls int32
	}{
		{name: "temporary DNS failure", failure: &net.DNSError{Err: "temporary", IsTemporary: true}, wantCalls: 2},
		{name: "permanent failure", failure: errors.New("bad certificate"), wantCalls: 1},
	} {
		t.Run(tt.name, func(t *testing.T) {
			var calls atomic.Int32
			hc := &http.Client{Transport: roundTripFunc(func(r *http.Request) (*http.Response, error) {
				if calls.Add(1) == 1 {
					return nil, tt.failure
				}
				return &http.Response{StatusCode: http.StatusOK, Body: io.NopCloser(strings.NewReader(`[]`)), Header: make(http.Header), Request: r}, nil
			})}
			c, err := New(WithHTTPClient(hc), WithRetryPolicy(RetryPolicy{
				MaxAttempts: 2,
				BaseDelay:   time.Millisecond,
				MaxDelay:    time.Millisecond,
			}))
			if err != nil {
				t.Fatal(err)
			}
			_, err = c.TVL().GetProtocols(context.Background())
			if tt.wantCalls == 2 && err != nil {
				t.Fatal(err)
			}
			if tt.wantCalls == 1 && err == nil {
				t.Fatal("expected the permanent transport error")
			}
			if got := calls.Load(); got != tt.wantCalls {
				t.Errorf("requests = %d, want %d", got, tt.wantCalls)
			}
		})
	}
}

func TestNoRetryOn404(t *testing.T) {
	_, c, log := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	}, WithRetryPolicy(RetryPolicy{MaxAttempts: 3, BaseDelay: time.Millisecond}))
	var nf *NotFoundError
	if _, err := c.TVL().GetProtocol(context.Background(), "x"); !errors.As(err, &nf) {
		t.Fatal("expected NotFoundError")
	}
	if len(log.all()) != 1 {
		t.Errorf("requests = %d, want 1", len(log.all()))
	}
}

func TestContextCancellation(t *testing.T) {
	_, c, _ := testServer(t, func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
		okJSON(w, r)
	})
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()
	_, err := c.TVL().GetProtocols(ctx)
	var te *TransportError
	if !errors.As(err, &te) {
		t.Fatalf("err = %v (%T), want TransportError", err, err)
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("err should wrap context.DeadlineExceeded: %v", err)
	}
}

func TestTimeoutConfig(t *testing.T) {
	if _, err := New(WithTimeout(0)); err == nil {
		t.Error("expected config error for timeout=0")
	}
	var ce *ConfigError
	if _, err := New(WithTimeout(-time.Second)); !errors.As(err, &ce) {
		t.Errorf("err = %v, want ConfigError", err)
	}
}

func TestNilOptionReturnsConfigError(t *testing.T) {
	var opt Option
	_, err := New(opt)
	var configErr *ConfigError
	if !errors.As(err, &configErr) {
		t.Fatalf("err = %v (%T), want ConfigError", err, err)
	}
}

func TestBaseURLsAreCopied(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(okJSON))
	t.Cleanup(srv.Close)
	baseURLs := map[string]string{"https://api.llama.fi": srv.URL}
	c, err := New(WithBaseURLsForTesting(baseURLs))
	if err != nil {
		t.Fatal(err)
	}
	baseURLs["https://api.llama.fi"] = "http://127.0.0.1:1"

	if _, err := c.TVL().GetProtocols(context.Background()); err != nil {
		t.Fatalf("client was changed by a caller map mutation: %v", err)
	}
}

func TestConcurrentUse(t *testing.T) {
	_, c, _ := testServer(t, okJSON)
	var wg sync.WaitGroup
	for i := 0; i < 16; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := c.TVL().GetProtocols(context.Background()); err != nil {
				t.Error(err)
			}
			if _, err := c.Yields().GetPools(context.Background()); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (f roundTripFunc) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}
