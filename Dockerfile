FROM golang:1.25.9-alpine AS builder

WORKDIR /app
COPY go.mod ./
COPY main.go ./

ARG TARGETOS=linux
ARG TARGETARCH=amd64
RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -ldflags="-s -w" -o /out/whatismyip .

FROM scratch
COPY --from=builder /out/whatismyip /whatismyip

EXPOSE 8080
ENTRYPOINT ["/whatismyip"]
