# Stage 1: Build
FROM python:3.12 as builder
WORKDIR /src
COPY app.py .

# Stage 2: Final minimal image
FROM python:3-12-alpine
WORKDIR /app
COPY --from=builder /src/app.py .
CMD ["python", "app.py"]