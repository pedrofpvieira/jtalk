FROM golang:alpine
LABEL maintainer="Tie Fighters <tiefighters.mall.ptc@jumia.com>"

# setup dir
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

EXPOSE 7070