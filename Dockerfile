FROM golang:1.7.1-alpine

RUN apk add --update git && apk add --update make && rm -rf /var/cache/apk/*

ADD . /go/src/github.com/ernestio/generic-adapter
WORKDIR /go/src/github.com/ernestio/generic-adapter

RUN make deps && go install

ENTRYPOINT ./entrypoint.sh
