# Go IP Metadata Service

Small Go webservice that returns request metadata as JSON:

- `ip`
- `user-agent`
- `x-forwarded-for`

## Run locally

```bash
go run .
```

Server listens on `8080` by default.
Set a custom port with `PORT`, for example:

```bash
PORT=9090 go run .
```

## Test

```bash
go test ./...
```

## Docker

Build image:

```bash
docker build -t whatismyip .
```

Run container:

```bash
docker run --rm -p 8080:8080 whatismyip
```

## CI and Security Pipeline

GitHub Actions runs two workflows on pushes and pull requests to `main`:

- `CI` workflow:
  - `go test ./...`
  - `go build ./...`
  - `govulncheck ./...`
  - Trivy filesystem scan (`trivy fs`)
  - Docker build and Trivy image scan (`trivy image`)
- `CodeQL` workflow:
  - GitHub CodeQL analysis for Go
  - Weekly scheduled scan

Security gates:

- Trivy scans fail the workflow when `HIGH` or `CRITICAL` vulnerabilities are found.
- `govulncheck` fails the workflow on detected Go vulnerabilities.
- CodeQL findings appear in the repository's code scanning alerts.

## Example response

```json
{
  "ip": "127.0.0.1",
  "user-agent": "curl/8.7.1",
  "x-forwarded-for": ""
}
```
