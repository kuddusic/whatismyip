FROM golang:1.25.9-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o /out/whatismyip .

FROM scratch
COPY --from=builder /out/whatismyip /whatismyip

EXPOSE 8080
ENTRYPOINT ["/whatismyip"]
