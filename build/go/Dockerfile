FROM golang:1.19.1-alpine
RUN apk update && apk add git
RUN mkdir /go/src/api

WORKDIR /go/src/api

ADD . /go/src/api

RUN go install github.com/cosmtrek/air@latest

RUN touch .air.toml

CMD [ "air","-c",".air.toml" ]