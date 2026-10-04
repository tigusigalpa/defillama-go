# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add `WithMaxResponseBodyBytes`, with a 32 MiB default limit, and the
  inspectable `ResponseBodyTooLargeError` / `ErrResponseBodyTooLarge` pair.
  The limit applies consistently to typed decoding and receipt capture.

### Changed

- Finalize every original HTTP response body before returning, observing, or
  retrying: it is read within the configured limit, boundedly drained, and
  closed exactly once. Receipts now expose `CompletedAt` and `Complete`; an
  incomplete receipt retains only the safely captured body prefix and digest.
- Preserve API, decode, read, drain, and close failures together using Go error
  wrapping, so each cause remains available through `errors.Is` and
  `errors.As`. Response-limit failures are not retried.
- Set the exported package `Version` and default User-Agent to the `v1.2.1`
  release value.

## [1.1.0] - 2026-10-04

### Added

- Add lossless receipts for the provider-native protocol and yield routes:
  `GetProtocolsReceipt`, `GetProtocolReceipt`, `GetPoolsReceipt`, and
  `GetPoolChartReceipt`. Receipts preserve immutable response bytes, redacted
  route provenance, capture time, and a SHA-256 digest; `Decode` uses
  `json.Decoder.UseNumber`.
- Add open, selected-field lossless DTOs for protocol, protocol detail, yield
  pool, and yield chart responses. They preserve exact number lexemes and the
  distinction between absent, null, and zero without claiming a closed provider
  schema.
- Add `WithReceiptObserver` for applications that need one receipt per actual
  HTTP response, including responses that cause a retry.

### Security

- Redact the Pro API key from HTTP error response headers and bodies, as well
  as from nested `url.Error` values returned by the standard HTTP client.
- Prevent Pro requests from following redirects to another origin, which could
  expose the key embedded in the request URL.
- Redact the key when request construction fails before an HTTP call is sent.

- Reject an empty token list before it can produce an invalid price-route URL.
- Copy test base-URL overrides at client construction, so later caller-side
  map changes cannot mutate a live client or introduce a data race.
- Return `ConfigError` for a nil functional option instead of panicking.
- Reject empty required query values and invalid path enum values locally.
- Reject a successful response containing trailing garbage or multiple JSON
  values with `DecodeError`.
- Retry transient network errors only; permanent transport failures are
  returned immediately even when retries are enabled.
- Replace the single mixed Free/Pro example with focused Free and Pro examples,
  and expand the README with runnable setup, routing, query, retry, and error
  handling guidance.

## [0.1.0] - 2026-09-23

### Added

- Initial release covering all 132 GET operations of the DefiLlama API
  (31 Free + 101 Pro), generated from the pinned `spec/defillama-api.json`.
- `New()` client with 21 concurrency-safe services: TVL, Prices, Stablecoins,
  Yields, Volumes, Fees, Emissions, Ecosystem, Bridges, ETFs, Narratives,
  Account, DAT, Treasury, Oracles, Forks, Dimensions, FinancialStatements,
  Equities, PreIPO, RWA.
- Central `routeRegistry` generated from the spec with per-operation origin,
  tier, Pro override path, parameter metadata and official docs URL.
- Pro authentication via URL path segment with full key redaction in errors.
- `WithPreferProForFree` applying the official Free→Pro mapping.
- Functional options: `WithAPIKey`, `WithHTTPClient`, `WithTimeout`,
  `WithRetryPolicy`, `WithUserAgent`, `WithPreferProForFree`,
  `WithBaseURLsForTesting`.
- Typed errors: `ConfigError`, `ProAPIKeyRequiredError`, `NotFoundError`,
  `RateLimitError` (with `RetryAfter`), `APIError`, `TransportError`,
  `DecodeError` — all compatible with `errors.Is`/`errors.As`.
- Typed models (`Protocol`, `ProtocolDetails`, `Chain`, `CoinPrice`,
  `Stablecoin`, `YieldPool`, `APIUsage`, `EquityCompany`, `PreIPOCompany`,
  `RWAAsset`, `Block`) retaining the raw payload.
- Bounded retry policy: GET-only, network/429/5xx, exponential backoff with
  jitter, `Retry-After` support, context-aware.
- Data-driven contract test asserting registry/spec/api-index parity (132 ops).
- Documentation index: `docs/api-index.md` (132 official doc links).
