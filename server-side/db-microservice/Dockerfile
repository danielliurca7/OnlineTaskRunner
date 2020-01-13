FROM golang:1.12.7-alpine3.10

RUN mkdir /app
ADD ./db-microservice /app
ADD ./utils /utils
ADD ./data_structures /data_structures

WORKDIR /app

RUN apk add git
RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/rs/xid
RUN go build -o db-main .

EXPOSE 10000
CMD ["/app/db-main"]