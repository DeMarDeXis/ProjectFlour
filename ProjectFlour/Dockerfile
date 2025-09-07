FROM golang:1.24-bullseye AS builder
ENV GOMAXPROCS=1

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && \
    rm -rf /go/pkg/mod/cache/download/*.zip

COPY . .

RUN CGO_ENABLED=0 GOOS=linux \
    go build -ldflags="-s -w" -o /app/main ./cmd/pFlour && \
    rm -rf /go/pkg/mod /root/.cache/go-build

FROM alpine:3.18

RUN apk --no-cache add ca-certificates && \
    adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/main /app/main
COPY --from=builder /app/config/prod.yaml /app/config/prod.yaml

RUN chown -R appuser:appuser /app

USER appuser

EXPOSE 8080 8081
CMD ["./main", "--config=/app/config/prod.yaml"]