FROM golang:latest

RUN mkdir -p /go/src/urlshortener

ADD . /go/src/urlshortener

RUN apt-get update && apt-get install -y xdg-utils && apt-get install unzip

WORKDIR /go/src/urlshortener

RUN go get
RUN go get github.com/tockins/realize

EXPOSE 8081

CMD realize start --run
