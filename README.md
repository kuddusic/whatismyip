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

## Example response

```json
{
  "ip": "127.0.0.1",
  "user-agent": "curl/8.7.1",
  "x-forwarded-for": ""
}
```
