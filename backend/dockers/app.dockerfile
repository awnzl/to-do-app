FROM golang:1.24-alpine AS builder

WORKDIR /build

COPY . .
RUN go mod download
RUN go build -o server ./cmd/server


FROM alpine:latest
RUN apk add --no-cache curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.17.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv migrate /usr/local/bin/migrate && \
    apk del curl

WORKDIR /app
COPY --from=builder /build/server .
COPY --from=builder /build/migrations ./migrations

RUN chmod +x /app/server

EXPOSE 8080
CMD ["/app/server"]
