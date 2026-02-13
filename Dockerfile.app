FROM golang:1.24-alpine AS builder

WORKDIR /src
COPY app/go.mod app/go.sum ./
RUN go mod download

COPY app/ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/ilicense-lite .

FROM alpine:3.21

WORKDIR /app
RUN addgroup -S app && adduser -S -G app app \
  && mkdir -p /app/etc /app/logs \
  && chown -R app:app /app

COPY --from=builder /out/ilicense-lite /app/ilicense-lite
COPY app/etc/app.docker.yaml /app/etc/app.yaml

USER app
EXPOSE 8080

ENTRYPOINT ["/app/ilicense-lite", "-config", "/app/etc/app.yaml"]
