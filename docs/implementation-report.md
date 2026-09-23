# Implementation report — defillama-go 0.1.0

Generated 2026-09-23 on Windows / Go 1.26.4.

## Coverage

- **132 GET operations** registered in `routes.go` — 31 Free + 101 Pro.
- **132 public service methods** across 21 services, verified by
  `TestEveryRouteHasPublicServiceMethod` (reflection over all service types).
- **132 unique official documentation URLs** in `docs/api-index.md`,
  asserted in CI.
- Route entries carry: HTTP method, path, tier, per-operation server, Pro
  override path (all 31 Free routes), service, method name, docs URL and the
  full parameter contract (name/in/required/type/enum).

## Commands executed and results

| Command | Result |
|---|---|
| `go mod tidy` | ✅ clean, zero dependencies |
| `go build ./...` | ✅ |
| `go vet ./...` | ✅ |
| `staticcheck ./...` | ✅ (2026.2.1) |
| `go test ./...` | ✅ `ok github.com/tigusigalpa/defillama-go` |
| `go test -race ./...` | ✅ `ok` |
| `gofmt -l .` | ✅ empty |
| `golangci-lint run` | ⚠️ not installed locally; configured in CI (`.golangci.yml` + GitHub Action) |

## Verified behaviors

- Per-operation Free origins; Pro URL `https://pro-api.llama.fi/{key}/<path>`.
- `ProAPIKeyRequiredError` before any request is built (zero requests recorded).
- `WithPreferProForFree` rewrites all 31 Free routes to official Pro paths.
- Coins list: each identifier path-escaped once, joined by `,`.
- `/batchHistorical` JSON-encodes the `coins` map into a single query value.
- Boolean query params serialize as lowercase `true`/`false`; nil fields omitted.
- Required query flags enforced (`excludeTotalDataChart` etc.); spec enums
  validated locally (`range`, `key`, `currency`, `timeframe`).
- 404 → `NotFoundError`; 429 → `RateLimitError` with `RetryAfter`; 5xx →
  `APIError` (status, redacted URL, headers, ≤4096-byte body); malformed JSON →
  `DecodeError`; network failure → `TransportError` wrapping the cause
  (`errors.Is`/`errors.As` compatible).
- API key redaction asserted: absent from `%+v` output and `APIError.URL`.
- Retry: 503 → success on attempt 2; 404 never retried; context cancellation
  surfaces as `TransportError` wrapping `context.DeadlineExceeded`.
- Concurrency: 16 goroutines across services under `-race` — clean.
- Live smoke check (manual, Free API only): `go run ./examples/basic` fetched
  8,332 protocols and the current BTC price.

## Honest limitations

- `GET /usage/APIKEY` uses the literal `APIKEY` path segment per the pinned
  spec (the key is already in the Pro prefix).
- Dynamic/evolving responses return `map[string]any` / `[]any`; only stable
  shapes have typed structs (all with `Raw` payload retention).
- `WithTimeout` applies to the default `*http.Client`; a custom
  `WithHTTPClient` manages its own timeouts.
- Retries are opt-in (`WithRetryPolicy`, default off).
- golangci-lint was not run locally (not installed); CI runs it on every PR.

## Maintenance review — 2026-09-23

This follow-up review hardened client behavior without changing the endpoint
surface or the pinned-spec coverage.

- Pro API keys are now redacted not only from request URLs, but also from HTTP
  error headers, error bodies, nested `url.Error` diagnostics, and request
  construction errors. Pro redirects are limited to the original origin.
- Empty token lists now fail locally instead of generating an invalid price URL.
- Empty required query values and invalid path enum values fail locally;
  successful responses with trailing invalid JSON report `DecodeError`.
- Retries now distinguish transient network errors from permanent transport
  failures; cancellation stops retry handling immediately.
- Test base-URL maps are copied at `New` time, preventing caller mutation from
  changing a live client or racing with requests; a nil functional option now
  returns `ConfigError` rather than panicking.
- The Free example now has bounded context handling and guards against empty or
  incomplete responses. A separate Pro example demonstrates key setup, retry,
  typed rate-limit handling, usage, and a Pro chart request.

The updated implementation was checked locally with:

| Command | Result |
|---|---|
| `go mod tidy` | ✅ clean, zero dependencies |
| `go test ./...` | ✅ package and both examples compile |
| `go test -race ./...` | ✅ clean |
| `go vet ./...` | ✅ clean |
| `staticcheck ./...` | ✅ clean |
| `gofmt -l .` | ✅ empty |
| `golangci-lint run` | ⚠️ unavailable locally; still configured in CI |
