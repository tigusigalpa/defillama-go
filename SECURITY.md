# Security Policy

## Supported versions

| Version | Supported |
|---------|-----------|
| 0.1.x   | ✅        |

## Reporting a vulnerability

Please **do not** open a public issue for security vulnerabilities.

Report them privately to **sovletig@gmail.com** with:

- a description of the issue and its impact,
- steps to reproduce or a proof of concept,
- affected versions.

You should receive an acknowledgement within a few days.

## Notes on API key handling

- The DefiLlama Pro API key is embedded as a single URL path segment
  (`https://pro-api.llama.fi/{API_KEY}/...`) — this is the API's own
  authentication scheme, not a choice of this SDK.
- The SDK redacts the key from every exception message, stored URL and log
  line. If you find a path where the key can leak, treat it as a security bug.
- Never commit real keys; use the `DEFILLAMA_API_KEY` environment variable.
