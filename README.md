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
docker build --platform linux/amd64 -t whatismyip:amd64 .
```

Run container:

```bash
docker run --rm -p 8080:8080 whatismyip:amd64
```

Build and push multi-arch image (`linux/amd64` + `linux/arm64`) with Buildx:

```bash
docker buildx build \
  --platform linux/amd64,linux/arm64 \
  -t ghcr.io/<owner>/whatismyip:latest \
  --push .
```

Inspect published manifest platforms:

```bash
docker buildx imagetools inspect ghcr.io/<owner>/whatismyip:latest
```

## CI and Security Pipeline

GitHub Actions runs three workflows:

- `CI` workflow:
  - Triggers on pushes and pull requests to `main`
  - Also supports manual `workflow_dispatch` runs
  - Manual runs include `use_self_hosted` input to target a self-hosted runner
  - `go test ./...`
  - `go build ./...`
  - `govulncheck ./...`
  - Trivy filesystem scan (`trivy fs`)
  - Docker Buildx + Trivy image scan (`trivy image`) for:
    - `linux/amd64`
    - `linux/arm64`
- `CodeQL` workflow:
  - GitHub CodeQL analysis for Go
  - Supports manual `workflow_dispatch` runs with `use_self_hosted` input
  - Weekly scheduled scan
- `Publish Container Image` workflow:
  - Publishes multi-arch container images to `ghcr.io`
  - Triggers on version tags (`v*`) and manual dispatch
  - Manual runs support optional `image_tag` and `use_self_hosted` inputs

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
