FROM golang:1.12.7-alpine3.10

RUN mkdir /app
ADD ./client-microservice /app
ADD ./utils /utils
ADD ./data_structures /data_structures

WORKDIR /app

RUN apk add git
RUN go get github.com/gorilla/mux
RUN go get github.com/graarh/golang-socketio
RUN go build -o client-main .

EXPOSE 8000
CMD ["/app/client-main"]