FROM golang:1.10-alpine3.7 AS builder
RUN apk add -U git
RUN go get -u -v github.com/wzshiming/pic2ascii/cmd/pic2ascii

FROM alpine:3.7
COPY --from=builder $GOBIN/pic2ascii /usr/local/bin/pic2ascii