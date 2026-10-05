# DefiLlama Golang Client/SDK/Library

![DefiLlama Go Golang SDK Client](https://i.postimg.cc/pr5KwwKP/defillama-golang-hero-github.jpg)

[![CI](https://github.com/tigusigalpa/defillama-go/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/defillama-go/actions/workflows/ci.yml)
[![Tests](https://github.com/tigusigalpa/defillama-go/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/defillama-go/actions/workflows/test.yml)
[![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?style=flat-square&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)
[![CodeQL](https://github.com/tigusigalpa/defillama-go/actions/workflows/codeql.yml/badge.svg?branch=main)](https://github.com/tigusigalpa/defillama-go/actions/workflows/codeql.yml)
[![Codecov](https://codecov.io/gh/tigusigalpa/defillama-go/graph/badge.svg)](https://codecov.io/gh/tigusigalpa/defillama-go)
[![GitHub Release](https://img.shields.io/github/v/release/tigusigalpa/defillama-go?style=flat-square)](https://github.com/tigusigalpa/defillama-go/releases)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue?style=flat-square&logo=go)](https://pkg.go.dev/github.com/tigusigalpa/defillama-go)

`defillama-go` is a small, idiomatic Go client for the [DefiLlama Free and Pro APIs](https://defillama.com/). It gives an application one consistent, context-aware interface for protocol TVL, token prices, yields, stablecoins, fees, bridges, RWA, equities, and the rest of the documented API surface.

> [!TIP]
> 📚 **Explore the full documentation:** [DefiLlama Go Wiki →](https://github.com/tigusigalpa/defillama-go/wiki)

- All **132 GET operations** in the included OpenAPI snapshot are available: 31 Free and 101 Pro.
- The runtime uses only Go's standard library. Supply your own `*http.Client` when you need a proxy, custom transport, tracing, or different timeout policy.
- `New` performs no network request, every network method accepts `context.Context`, and one client is safe to share between goroutines.
- The route registry is generated from `spec/defillama-api.json`. The complete operation-to-method map is in [docs/api-index.md](docs/api-index.md).

This package is an API transport, not financial advice. DefiLlama metrics and token prices follow DefiLlama's methodology; use a decimal type in application code where exact monetary precision matters.

## Install

Requires Go 1.22 or newer.

```bash
go get github.com/tigusigalpa/defillama-go
```

## Your first request: no API key needed

Free endpoints work immediately. The following program reads the protocol list and Bitcoin's current USD price. In production, use a request-scoped context with a deadline just as shown here.

```go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	defillama "github.com/tigusigalpa/defillama-go"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := defillama.New(defillama.WithTimeout(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}

	protocols, err := client.TVL().GetProtocols(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if len(protocols) > 0 {
		fmt.Printf("%s TVL: $%.0f\n", protocols[0].Name, protocols[0].TVL)
	}

	prices, err := client.Prices().GetCurrentPrices(ctx, []string{
		"coingecko:bitcoin",
		"ethereum:0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",
	})
	if err != nil {
		log.Fatal(err)
	}
	btc, ok := prices["coingecko:bitcoin"]
	if !ok {
		log.Fatal("DefiLlama response did not include bitcoin")
	}
	fmt.Printf("BTC: $%.2f\n", btc.Price)
}
```

You can run the equivalent checked example from this repository:

```bash
go run ./examples/basic
```

Free operations use the origin specified for that individual route, rather than forcing every call through one host. Depending on the operation, that is `api.llama.fi`, `coins.llama.fi`, `stablecoins.llama.fi`, or `yields.llama.fi`.

## Pro API

Create a key in DefiLlama, keep it outside source control, and provide it to the client. For example:

```bash
# macOS / Linux
export DEFILLAMA_API_KEY='your-key'
go run ./examples/pro

# PowerShell
$env:DEFILLAMA_API_KEY = 'your-key'
go run ./examples/pro
```

```go
client, err := defillama.New(
	defillama.WithAPIKey(os.Getenv("DEFILLAMA_API_KEY")),
	defillama.WithTimeout(10*time.Second),
	defillama.WithRetryPolicy(defillama.RetryPolicy{
		MaxAttempts: 3,
		BaseDelay:  250 * time.Millisecond,
		MaxDelay:   2 * time.Second,
		Jitter:     true,
	}),
)
if err != nil {
	return err
}

usage, err := client.Account().GetUsage(ctx)
if err != nil {
	return err
}
if credits, ok := usage.CreditsLeft(); ok {
	fmt.Printf("credits remaining: %.0f\n", credits)
}

chart, err := client.TVL().GetProtocolTVLChart(ctx, "aave", nil)
if err != nil {
	return err
}
fmt.Println("chart points:", len(chart))
```

DefiLlama authenticates Pro requests with the key in the URL path:

```
https://pro-api.llama.fi/{API_KEY}/<endpoint path>
```

It is deliberately never sent as an HTTP header or query parameter. The client escapes the key as one path segment and redacts it from URLs, response diagnostics, and nested HTTP errors. Calling a Pro method without a key returns `*defillama.ProAPIKeyRequiredError` before the SDK builds or sends a request.

Pro requests can follow redirects on the same origin. The SDK returns a redirect to another origin as an `APIError` with the redirect status, without following it.

If you already have a Pro key and want the Free methods to use DefiLlama's official Pro equivalents too, opt in explicitly:

```go
client, err := defillama.New(
	defillama.WithAPIKey(os.Getenv("DEFILLAMA_API_KEY")),
	defillama.WithPreferProForFree(true),
)
```

This only changes routing for the 31 Free endpoints that have an official Pro mapping. It is off by default, so a configured key never silently changes the default Free routing.

## Common tasks

The snippets in this section assume you already have a `client` and a `ctx`,
as in the first example. They fit inside a function that returns an `error`.
Examples marked **Pro** need a client configured with `WithAPIKey`.

### Explore a protocol's TVL (Free)

The protocol detail endpoint has a flexible response shape. Start with the
typed accessors, then inspect `Raw` when you need a field the SDK does not
model yet:

```go
protocol, err := client.TVL().GetProtocol(ctx, "aave")
if err != nil {
	return err
}
fmt.Printf("%s (%s): %d TVL samples\n",
	protocol.Name(), protocol.Slug(), len(protocol.TVLHistory()))

if ethereum, ok := protocol.ChainTVLs()["Ethereum"]; ok {
	fmt.Printf("Ethereum breakdown: %v\n", ethereum)
}
if description, ok := protocol.Raw["description"].(string); ok {
	fmt.Println(description)
}
```

`ChainTVLs()` returns raw per-chain values rather than assuming every chain has
the same shape. Check types before using fields from `Raw` in calculations.

### Ask for a historical token price

Use DefiLlama token identifiers in the `chain:address` form; CoinGecko assets use `coingecko:<id>`.

```go
at := time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
prices, err := client.Prices().GetHistoricalPrices(ctx,
	[]string{"coingecko:ethereum", "ethereum:0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"},
	at,
)
if err != nil {
	return err
}
eth, ok := prices["coingecko:ethereum"]
if !ok {
	return fmt.Errorf("historical response did not include ethereum")
}
fmt.Println(eth.Price)
```

### List stablecoins with prices (Free)

`GetStablecoins` returns typed records. Request prices explicitly, and keep
the output small when you're exploring a large list:

```go
stablecoins, err := client.Stablecoins().GetStablecoins(ctx,
	&defillama.StablecoinsListOptions{
		IncludePrices: defillama.Ptr(true),
	},
)
if err != nil {
	return err
}
for i, coin := range stablecoins {
	if i >= 3 {
		break
	}
	fmt.Printf("%s (%s): %d chains, price %.4f\n",
		coin.Name, coin.Symbol, len(coin.Chains), coin.Price)
}
```

When an optional field's presence matters, inspect `coin.Raw` instead of
treating a zero-valued Go field as proof that the API returned zero.

### Filter earn pools

Option structs make optional values explicit. `defillama.Ptr` keeps examples tidy, and `nil` fields are not serialized.

```go
result, err := client.Yields().QueryEarnPools(ctx, defillama.EarnPoolsQuery{
	Chain:      defillama.Ptr("Ethereum"),
	Stablecoin: defillama.Ptr(true),
	MinTVL:     defillama.Ptr(1_000_000.0),
	Page:       defillama.Ptr(int64(1)),
	Limit:      defillama.Ptr(int64(50)),
})
if err != nil {
	return err
}
fmt.Printf("filtered response fields: %d\n", len(result))
```

`QueryEarnPools` reflects the API's evolving response as `map[string]any`. For
the stable, typed Free pool list, use `client.Yields().GetPools(ctx)`, which
returns `[]defillama.YieldPool`.

### Request a price chart

```go
chart, err := client.Prices().GetChart(ctx,
	[]string{"coingecko:bitcoin"},
	&defillama.ChartOptions{
		Span:   defillama.Ptr(30.0), // 30 data points
		Period: defillama.Ptr("1d"),
	},
)
if err != nil {
	return err
}
fmt.Println("chart response fields:", len(chart))
```

The SDK validates declared enum values and required flags before making a request. For routes with required Boolean flags, such as DEX overview, use the value-based option struct:

```go
overview, err := client.Volumes().GetDEXOverview(ctx, defillama.OverviewOptions{
	ExcludeTotalDataChart:          true,
	ExcludeTotalDataChartBreakdown: true,
})
if err != nil {
	return err
}
_ = overview
```

### Inspect real-world assets (Pro)

RWA assets have typed fields for common data and a `Raw` map for additional
fields returned by the API:

```go
assets, err := client.RWA().GetCurrentAssets(ctx)
if err != nil {
	return err
}
for i, asset := range assets {
	if i >= 3 {
		break
	}
	fmt.Printf("%s (%s), issuer: %s\n", asset.AssetName, asset.Ticker, asset.Issuer)
}
if len(assets) > 0 {
	fmt.Printf("first asset's complete payload: %v\n", assets[0].Raw)
}
```

### Filter recent bridge transactions (Pro)

Bridge IDs come from `client.Bridges().GetBridges(ctx, nil)`. After choosing
one, pass Unix timestamps and a bounded result limit:

```go
bridgeID := int64(1) // Replace with an ID from GetBridges.
end := time.Now().UTC()
start := end.Add(-24 * time.Hour)

transactions, err := client.Bridges().GetBridgeTransactions(ctx, bridgeID,
	&defillama.BridgeTransactionsOptions{
		Starttimestamp: defillama.Ptr(start.Unix()),
		Endtimestamp:   defillama.Ptr(end.Unix()),
		Sourcechain:    defillama.Ptr("Ethereum"),
		Limit:          defillama.Ptr(int64(100)),
	},
)
if err != nil {
	return err
}
fmt.Printf("bridge response fields: %d\n", len(transactions))
```

The transaction response is `map[string]any`; inspect its fields before
converting them to application-specific types. `Limit` must be between 1 and
6000, and the SDK rejects values outside that range before sending a request.

## Handle errors deliberately

All service errors are typed and work with `errors.As`. URLs in errors are already safe to record in logs.

```go
_, err := client.TVL().GetProtocol(ctx, "not-a-real-protocol")

var (
	notFound  *defillama.NotFoundError
	rateLimit *defillama.RateLimitError
	apiErr    *defillama.APIError
	transport *defillama.TransportError
	decode    *defillama.DecodeError
	missingKey *defillama.ProAPIKeyRequiredError
)

switch {
case errors.As(err, &notFound):
	// The resource does not exist (HTTP 404).
case errors.As(err, &rateLimit):
	// HTTP 429. rateLimit.RetryAfter is zero if the server did not provide it.
	return fmt.Errorf("try again in %s: %w", rateLimit.RetryAfter, err)
case errors.As(err, &apiErr):
	// Other non-2xx responses: StatusCode, Header, Body, and a redacted URL.
case errors.As(err, &transport):
	// DNS, TLS, timeout, or cancellation; errors.Is reaches the root cause.
case errors.As(err, &decode):
	// The server returned a successful response that was not valid JSON.
case errors.As(err, &missingKey):
	// A Pro method was called without WithAPIKey.
case err != nil:
	return err
}
```

Retries are disabled by default. When configured, only GET requests with transient transport errors, HTTP 429, or HTTP 5xx are retried. Backoff is capped, includes optional jitter, honors `Retry-After`, and stops promptly when the context is cancelled or reaches its deadline. Decode errors and ordinary 4xx responses are never retried.

## Preserve exact protocol and yield receipts

The ordinary service methods are the convenient compatibility interface: their
documented numeric fields use `float64`, and unknown fields in `map[string]any`
follow Go's usual JSON decoding rules. That is not a lossless boundary for
economic values.

For provider-native protocol and yield data, use the receipt methods. A
complete receipt preserves the exact response bytes, route provenance, redacted
request URL, attempt number, and SDK capture time. `Decode` uses
`json.Decoder.UseNumber`, so an untyped number remains a `json.Number` with its
original lexeme. This is especially useful when a caller owns an immutable
ingestion receipt or needs to distinguish an absent member, `null`, and numeric
zero.

```go
receipt, err := client.TVL().GetProtocolReceipt(ctx, "aave")
if err != nil {
	return err
}

var details defillama.LosslessProtocolDetails
if err := receipt.Decode(&details); err != nil {
	return err
}
name, _, err := details.Name()
if err != nil {
	return err
}
history, present, err := details.TVLHistory()
if err != nil {
	return err
}
if !present || len(history) == 0 {
	return fmt.Errorf("protocol history unavailable")
}
date, present := history[0].Field("date")
if !present {
	return fmt.Errorf("provider history date unavailable")
}
fmt.Printf("%s history date JSON: %s\n", name, date.RawJSON())
fmt.Println("receipt SHA-256:", receipt.BodySHA256())
```

`CapturedAt` is recorded as soon as the HTTP response is received;
`CompletedAt` is recorded only after the original response body has been read,
boundedly drained, and closed. Always check `Complete` before treating `Body()`
as a whole response. If a read, limit, drain, or close step fails, the receipt
still exposes the immutable prefix and its SHA-256 digest for diagnostics, but
`Complete` is false and the returned error keeps every underlying cause
discoverable through `errors.Is` and `errors.As`.

Every response body, including ordinary typed calls, is limited to 32 MiB by
default. Choose a smaller or larger positive limit when constructing the client;
an over-limit response is never retried and matches
`defillama.ErrResponseBodyTooLarge`.

```go
client, err := defillama.New(
	defillama.WithMaxResponseBodyBytes(8<<20), // 8 MiB
)
if err != nil {
	return err
}

receipt, err := client.TVL().GetProtocolsReceipt(ctx)
if err != nil {
	if errors.Is(err, defillama.ErrResponseBodyTooLarge) {
		// receipt, when non-nil, holds the safely bounded prefix.
		return fmt.Errorf("protocol response exceeds the ingestion budget: %w", err)
	}
	return err
}
if !receipt.Complete {
	return fmt.Errorf("incomplete protocol receipt")
}
```

The selected receipt methods are `TVL().GetProtocolsReceipt`,
`TVL().GetProtocolReceipt`, `Yields().GetPoolsReceipt`, and
`Yields().GetPoolChartReceipt`. Decode their bodies into `LosslessProtocols`,
`LosslessProtocolDetails`, `LosslessYieldPools`, or `LosslessYieldChart`.
Those DTOs intentionally expose only selected convenience fields and retain all
other provider fields through open `Fields()` objects; they do not claim a
closed or complete provider schema.

Pool IDs and protocol IDs are provider-native identifiers, not universal asset
identities. Protocol history `date` values are provider Unix seconds, while a
yield chart's `timestamp` is provider ISO-8601 text. Both are distinct from
`receipt.CapturedAt`. Yield `apy`, `apyBase`, and `apyReward` remain separate
provider percentages, and `predictions` stays an unclassified provider output.

To retain every HTTP response, including responses that trigger a retry, pass
an observer when constructing the client. The callback runs in the calling
goroutine only after that response's original body is finalized; a
transport-only failure produces no receipt. Avoid logging or persisting receipt
bodies unless that is permitted by your application's data policy.

```go
client, err := defillama.New(
	defillama.WithReceiptObserver(func(receipt defillama.ResponseReceipt) {
		fmt.Printf("attempt %d: HTTP %d, complete=%t, %s\n",
			receipt.Attempt, receipt.StatusCode, receipt.Complete, receipt.BodySHA256())
	}),
)
if err != nil {
	return err
}
```

### Inspect every attempt without retaining receipts

When retry diagnostics matter, attach an observer to the context for one
operation. It receives an event for each actual attempt, including a
response-less transport failure. `OperationID` groups events if an application
uses a shared callback; `Attempt` is the one-based number within that operation.
This is separate from `WithReceiptObserver`: the latter remains response-only,
so no failed dial or cancelled request is turned into a synthetic receipt.

```go
ctx = defillama.WithAttemptObserver(ctx, func(attempt defillama.AttemptDiagnostic) {
	if !attempt.ResponseReceived {
		// A dial, timeout, or cancellation failed before an HTTP response.
		if errors.Is(attempt.Err, context.DeadlineExceeded) {
			fmt.Printf("attempt %d timed out\n", attempt.Attempt)
		}
		return
	}

	fmt.Printf("attempt %d: HTTP %d, complete=%t, outcome=%s\n",
		attempt.Attempt, attempt.StatusCode, attempt.ResponseComplete, attempt.Outcome)
})

protocols, err := client.TVL().GetProtocols(ctx)
if err != nil {
	return err // The terminal error remains the ordinary service error.
}
_ = protocols
```

The callback runs synchronously only after a received response body has been
finalized. It receives no headers or response bytes; `TargetURL` removes
user-info and redacts all query values. The SDK keeps no attempt buffer,
so create a fresh observer context per operation and retain only the diagnostic
data your application needs. `Err` preserves its original cause for
`errors.Is`/`errors.As`, but is never logged or serialized by the SDK.

## Services

Every service accessor returns a lightweight client-owned value. Keep the top-level `Client` and call the service that matches the data you need.

| Accessor | What it covers | A starting method |
|---|---|---|
| `TVL()` | Protocols, chains, TVL, Pro TVL metrics and charts | `GetProtocols(ctx)` |
| `Prices()` | Current/historical prices, charts, blocks | `GetCurrentPrices(ctx, coins)` |
| `Stablecoins()` | Stablecoin supply, charts, prices, dominance | `GetStablecoins(ctx, opts)` |
| `Yields()` | Free pools plus Pro earn, borrow, LST, and perps data | `GetPools(ctx)` |
| `Volumes()` | DEX, options, open interest, and derivatives volumes | `GetDEXOverview(ctx, opts)` |
| `Fees()` | Fees and revenue overview/summary data | `GetOverview(ctx, opts)` |
| `Emissions()` | Token unlocks | `GetEmissions(ctx)` |
| `Ecosystem()` | Categories, entities, forks, hacks, oracles, raises, treasuries | `GetHacks(ctx)` |
| `Bridges()` | Bridges, volumes, daily stats, transactions | `GetBridgeTransactions(ctx, id, opts)` |
| `ETFs()` | ETF snapshots and flows | `GetSnapshot(ctx)` |
| `Narratives()` | Narrative performance | `GetPerformance(ctx, period)` |
| `Account()` | Pro key usage | `GetUsage(ctx)` |
| `DAT()` | Digital Asset Treasury institutions | `GetInstitutions(ctx)` |
| `Treasury()` | Protocol treasury metrics and charts | `GetTreasuryChart(ctx, protocol, opts)` |
| `Oracles()` | Oracle metrics and charts | `GetOracleMetrics(ctx)` |
| `Forks()` | Fork metrics and charts | `GetForkMetrics(ctx)` |
| `Dimensions()` | Cross-metric chart and breakdown endpoints | `GetChart(ctx, metric, opts)` |
| `FinancialStatements()` | Protocol financial statements | `GetIncomeStatement(ctx, protocol)` |
| `Equities()` | Public company lists, statements, OHLCV, filings | `GetCompaniesList(ctx)` |
| `PreIPO()` | Pre-IPO companies, valuations, raises, integrations | `GetValuations(ctx, query)` |
| `RWA()` | Real-world assets, stats, and charts | `GetCurrentAssets(ctx)` |

For all 132 operations, parameter names, and links to the official reference, see [docs/api-index.md](docs/api-index.md). Stable response shapes use typed models such as `Protocol`, `Chain`, `CoinPrice`, `Stablecoin`, and `YieldPool`. Evolving API responses are intentionally returned as clearly named `map[string]any` or `[]any`; the typed models also retain their complete raw payload in `Raw`.

## Configuration

| Option | Default | Purpose |
|---|---|---|
| `WithAPIKey(key)` | none | Enables Pro endpoints. The SDK does not read environment variables itself. |
| `WithTimeout(d)` | 30 seconds | Timeout used by the default `*http.Client`. |
| `WithHTTPClient(client)` | SDK default | Inject a client for a proxy, tracing, or your own timeout policy. |
| `WithRetryPolicy(policy)` | one attempt | Enables bounded retries for retryable GET failures. |
| `WithUserAgent(value)` | `defillama-go/<version>` | Overrides the User-Agent header. |
| `WithPreferProForFree(true)` | `false` | Routes Free calls to their official Pro equivalents when a key is present. |
| `WithReceiptObserver(fn)` | none | Receives one immutable receipt for every HTTP response, including retry attempts. |
| `WithMaxResponseBodyBytes(n)` | 32 MiB | Sets a positive inclusive response-body limit for typed decoding and receipts; overflow returns a non-retryable `ResponseBodyTooLargeError`. |

`WithTimeout` does not alter a custom client supplied through `WithHTTPClient`; configure that client's timeouts yourself. `WithBaseURLsForTesting` exists for tests and advanced local mock-server setups, not ordinary production configuration.

For example, to use an HTTP client already configured by your application
(requires the `net/http` and `time` imports):

```go
httpClient := &http.Client{Timeout: 8 * time.Second}
client, err := defillama.New(defillama.WithHTTPClient(httpClient))
if err != nil {
	return err
}
// Reuse client for subsequent requests; it is safe to share across goroutines.
```

## Response precision and compatibility

The ordinary typed API decodes values supplied by DefiLlama as `float64`
transport values. Convert prices, TVL, and percentages to a decimal
representation before calculations where rounding matters. Stable typed models
retain unknown fields in `Raw`, but `Raw` is a compatibility map and does not
preserve raw number lexemes. Use the receipt APIs above when exact JSON
representation, field presence, or immutable source bytes matter.

The package follows semantic versioning. Review [CHANGELOG.md](CHANGELOG.md) when upgrading; endpoint coverage changes are visible in [docs/api-index.md](docs/api-index.md).

## Development

The automated tests use local HTTP servers; they neither contact DefiLlama nor need a key.

```bash
go test ./...
go test -race ./...
go vet ./...
gofmt -l .             # should print nothing
staticcheck ./...      # if installed
golangci-lint run      # if installed
```

See [CONTRIBUTING.md](CONTRIBUTING.md) for project conventions, [docs/release-checklist.md](docs/release-checklist.md) for release metadata checks, and [docs/implementation-report.md](docs/implementation-report.md) for the original coverage and verification record. To report a security issue, follow [SECURITY.md](SECURITY.md).

## License

MIT — © 2026 Igor Sazonov. See [LICENSE](LICENSE).
