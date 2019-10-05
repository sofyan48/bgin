## Builder
FROM golang:latest

RUN go get -u github.com/golang/dep/cmd/dep && \
    mkdir -p /go/src/app

WORKDIR /go/src/app
COPY    . .
ENV  GOPATH=/go
RUN     echo ${GOPATH} && \
        dep ensure && \
        dep ensure -add github.com/meongbego/go_boilerplate

EXPOSE 5000
