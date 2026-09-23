# Contributing

Thanks for your interest in improving `defillama-go`.

## Setup

```bash
git clone https://github.com/tigusigalpa/defillama-go.git
cd defillama-go
go mod tidy
```

## Checks (must pass before submitting a PR)

```bash
gofmt -l .            # must print nothing
go vet ./...
staticcheck ./...     # go install honnef.co/go/tools/cmd/staticcheck@latest
golangci-lint run     # if installed
go test ./...
go test -race ./...
```

Tests never hit the real API and never need an API key — they use
`httptest.Server` and `WithBaseURLsForTesting`.

## Architecture rules

- `spec/defillama-api.json` is the source of truth. Routes live in
  `routes.go` (generated); never hand-build base URLs in service methods.
- New endpoints: update the spec snapshot, regenerate `routes.go`, the service
  files and `docs/api-index.md`.
- Query parameters go into option structs in `options.go`; optional fields are
  pointers so `nil` is never serialized.
- Every exported identifier needs a GoDoc comment.
- Never log or expose the API key — errors carry only the redacted URL.
- Keep the package dependency-free (standard library only) unless a new
  dependency carries clear, obligatory value.

## Reporting issues

Open an issue at https://github.com/tigusigalpa/defillama-go/issues with a
minimal reproduction (redact any API keys).
