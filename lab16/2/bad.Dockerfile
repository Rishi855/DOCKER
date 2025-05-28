FROM node:18
RUN apt-get update
RUN apt-get install -y git
RUN npm install -g http-server