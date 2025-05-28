FROM node:18-alpine

# Create app directory and non-root user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

WORKDIR /app

COPY . .

RUN npm install

# Switch to non-root user
USER appuser

CMD ["node", "index.js"]
