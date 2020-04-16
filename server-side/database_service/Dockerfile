FROM golang:alpine

WORKDIR /app

ADD . .

RUN apk add git
RUN apk add wget

RUN go get github.com/gorilla/mux
RUN go get github.com/lib/pq
RUN go get go.mongodb.org/mongo-driver/mongo

RUN go build -o database_main .

CMD ["/app/database_main"]