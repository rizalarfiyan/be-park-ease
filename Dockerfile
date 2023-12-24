# Build stage
FROM golang:1.21.1-alpine AS builder
WORKDIR /app
COPY . .

# Build the Go application
RUN go mod download
RUN go mod verify
ENV GO111MODULE=on
RUN CGO_ENABLED=0 GOOS=linux go build -o be-park-ease
RUN go install github.com/pressly/goose/v3/cmd/goose@latest

# Final stage
FROM alpine:latest
WORKDIR /app

RUN apk add --no-cache tzdata && \
    rm -rf /var/cache/apk/*

# Copy the binary and other necessary files
COPY --from=builder /app/be-park-ease /app/
COPY --from=builder /app/database /app/database
COPY --from=builder /app/migrate.sh /app/.
COPY --from=builder /go/bin/goose /app/.

CMD ["/app/be-park-ease"]
