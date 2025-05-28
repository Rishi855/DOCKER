FROM nginx:alpine

HEALTHCHECK --interval=10s --timeout=3s \
    CMD wget --no-verbose --spider http://localhost || exit 1