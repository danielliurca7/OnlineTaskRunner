FROM golang:1.12.7-alpine3.10

WORKDIR /app

ADD ./compute-microservice /app
ADD ./utils /utils
ADD ./data_structures /data_structures

RUN apk add --update docker openrc
RUN rc-update add docker boot

RUN apk add git

RUN go get github.com/gorilla/mux
RUN go get github.com/graarh/golang-socketio
RUN go build -o compute-main .

EXPOSE 9000
CMD ["/app/compute-main"]