FROM golang:1.13.1-alpine

ENV GO111MODULE=on

RUN apk update && \
    apk add emacs curl git && \
    apk add build-base && \
    go get github.com/aws/aws-sdk-go && \
    go get github.com/pilu/fresh