FROM golang:1.13-alpine AS builder
MAINTAINER Nikolay Bondarenko <nikolay.bondarenko@protocol.one>

RUN apk add bash ca-certificates git

WORKDIR /application

ENV GO111MODULE=on

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /application/bin/1cgateway .

FROM alpine:3.9
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
WORKDIR /application
COPY --from=builder /application /application

ENTRYPOINT /application/bin/1cgateway
