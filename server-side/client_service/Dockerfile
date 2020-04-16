FROM golang:alpine

WORKDIR /app

ADD . .

RUN apk add git

RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/streadway/amqp
RUN go get github.com/graarh/golang-socketio

RUN go build -o client_main .

CMD ["/app/client_main"]