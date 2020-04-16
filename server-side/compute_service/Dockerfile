FROM golang:alpine

WORKDIR /app

ADD . .

RUN apk add --update docker openrc
RUN rc-update add docker boot

RUN apk add git

RUN go get github.com/gorilla/mux
RUN go get github.com/graarh/golang-socketio

RUN go build -o compute_main .

CMD ["/app/compute_main"]