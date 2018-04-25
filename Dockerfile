FROM golang:1.10-alpine3.7 AS builder
WORKDIR /go/src/github.com/wzshiming/pic2ascii/
COPY . .
RUN CGO_ENABLED=0 go install -a -ldflags '-s' ./cmd/...

FROM scratch
COPY --from=builder /go/bin/pic2ascii /
ENTRYPOINT [ "/pic2ascii" ]
