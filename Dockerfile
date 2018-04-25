FROM golang:1.10-alpine3.7 AS builder
WORKDIR /go/src/github.com/wzshiming/pic2ascii/
COPY . .
RUN go install ./cmd/...

FROM alpine:3.7
COPY --from=builder /go/bin/pic2ascii /usr/local/bin/
