FROM golang:1.21-alpine

RUN apk update && apk add git

WORKDIR /go/src/app