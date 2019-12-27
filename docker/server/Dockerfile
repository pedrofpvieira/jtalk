FROM golang:1.13-alpine
LABEL maintainer="Tie Fighters <tiefighters.mall.ptc@jumia.com>"

## TODO Convert this dockerfile to a multi-stage build with a build image, and a scratch image

## Add git
RUN apk update && apk upgrade && \
    apk add --no-cache git build-base

# Install staticcheck.io
RUN go get honnef.co/go/tools/cmd/staticcheck

# setup dir
RUN mkdir /app
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

EXPOSE 7070