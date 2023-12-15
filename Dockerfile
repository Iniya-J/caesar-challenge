FROM golang:alpine

WORKDIR /go/src/github.com/wesleyholiveira/caesar-challenge
COPY . .

RUN apk add git && \
    go build -o /go/bin/caesar

ENTRYPOINT ["/go/bin/caesar"]
