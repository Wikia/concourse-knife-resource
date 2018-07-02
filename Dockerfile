FROM golang:1.10.3-alpine

RUN apk add --no-cache git

COPY out /go/src/github.com/Wikia/concourse-knife-resource/out/.
COPY check /go/src/github.com/Wikia/concourse-knife-resource/check/.

WORKDIR /go/src/github.com/Wikia/concourse-knife-resource/check

RUN go get && \
    go build

COPY in /go/src/github.com/Wikia/concourse-knife-resource/in/.

WORKDIR /go/src/github.com/Wikia/concourse-knife-resource/in

RUN go get && \
    go build

FROM alpine:3.7

COPY --from=0 /go/src/github.com/Wikia/concourse-knife-resource/check/check /opt/resource/check
COPY --from=0 /go/src/github.com/Wikia/concourse-knife-resource/in/in /opt/resource/in
COPY --from=0 /go/src/github.com/Wikia/concourse-knife-resource/out/out /opt/resource/out
