# Stage 1: Build
FROM golang:alpine AS builder
WORKDIR /app
COPY main.go .
RUN go env -w GO111MODULE=off && go build -o hello

# Stage 2: Runtime
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/hello .
CMD ["./hello"]