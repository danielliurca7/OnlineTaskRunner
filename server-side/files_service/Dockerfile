FROM golang:alpine

WORKDIR /app

ADD . .

WORKDIR /app

RUN apk add git

RUN go get github.com/gorilla/mux
RUN go get github.com/go-redis/redis
RUN go get github.com/streadway/amqp

RUN go build -o files_main .

CMD ["/app/files_main"]