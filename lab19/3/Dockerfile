# Stage 1: Builder
FROM golang:1.20 AS builder
WORKDIR /src
COPY go.mod .
COPY main.go .
RUN CGO_ENABLED=0 go build -o app .

# Stage 2: Minimal image
FROM alpine
RUN adduser -D appuser
USER appuser
WORKDIR /app
COPY --from=builder /src/app .

CMD ["./app"]
