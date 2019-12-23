FROM node:alpine

# setup dir
RUN mkdir /app
WORKDIR /app

COPY client/package.json /app
COPY client/package-lock.json /app

RUN npm install

EXPOSE 3000