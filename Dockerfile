FROM golang:alpine AS builder
WORKDIR /go/src/github.com/wzshiming/pic2ascii/
COPY . .
RUN apk add -U gcc libc-dev ffmpeg-dev ffmpeg-libs
RUN go install -tags support_video -a -ldflags '-s -w' ./cmd/...

FROM alpine
COPY --from=builder /go/bin/pic2ascii /usr/local/bin/
RUN apk add -U --no-cache ffmpeg-libs
ENTRYPOINT [ "pic2ascii" ]
