FROM golang:1.17-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokend/bridge/core
COPY vendor .
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /usr/local/bin/core gitlab.com/tokend/bridge/core


FROM alpine:3.9

RUN apk add --no-cache ca-certificates
COPY --from=buildbase /usr/local/bin/core /usr/local/bin/core

ENTRYPOINT ["core"]
