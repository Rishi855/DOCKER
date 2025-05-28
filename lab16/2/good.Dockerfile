FROM node:18
RUN apt-get update && \
    apt-get install -y git && \
    npm install -g http-server && \
    apt-get clean && rm -rf /var/lib/apt/lists/*