FROM alpine
RUN apk add --no-cache curl

HEALTHCHECK --interval=5s --timeout=2s --retries=3 \
  CMD curl --fail http://localhost:9999 || exit 1

CMD ["sleep", "3600"]
