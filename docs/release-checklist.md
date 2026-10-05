# Release checklist

`client.go` is the single source of the SDK release version. Do not derive the
SDK version from `go.mod` or the bundled OpenAPI document: those files describe
the module language target and provider API snapshot, respectively.

## Before creating a release tag

- Update `Version` in `client.go` and verify that the default User-Agent is
  exactly `defillama-go/<Version>`.
- Add a matching `## [<Version>] - YYYY-MM-DD` entry to `CHANGELOG.md`.
- Run `go test ./...`, `go test -race ./...`, `go vet ./...`, and the configured
  `golangci-lint` version. The release-metadata test checks the version and
  changelog pairing without requiring a Git tag, so ordinary source builds and
  pull requests remain tag-independent.

## When publishing

- Create an annotated `v<Version>` tag from the verified commit.
- Confirm that the tag points at the intended commit:

  ```bash
  git tag --points-at HEAD
  ```

- Push the tag and confirm that the release workflow completed successfully.

The final tag comparison is deliberately a publishing check, not a normal build
or test prerequisite.
