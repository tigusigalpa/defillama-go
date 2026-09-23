# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Security

- Redact the Pro API key from HTTP error response headers and bodies, as well
  as from nested `url.Error` values returned by the standard HTTP client.
- Prevent Pro requests from following redirects to another origin, which could
  expose the key embedded in the request URL.
- Redact the key when request construction fails before an HTTP call is sent.

### Changed

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
