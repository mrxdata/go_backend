FROM golang:1.23.6-bullseye AS builder

LABEL maintainer="dv3"

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -ldflags="-s -w" -o apiserver .

FROM alpine:latest

RUN apk --no-cache add tini

COPY --from=builder ["/build/apiserver", "/build/resources/prk.der", "/"]

ENTRYPOINT ["/sbin/tini", "--", "/apiserver"]
