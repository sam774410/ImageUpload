FROM golang:latest

MAINTAINER "sam850118sam@gmail.com"

WORKDIR /app

RUN mkdir -p /app/upload/orign
RUN mkdir -p /app/upload/custom

ADD . /app


RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/disintegration/imaging


EXPOSE 3000

CMD ["go", "run", "main.go"]