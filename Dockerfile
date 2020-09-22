FROM golang:alpine

RUN apk update && apk add git

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o mygosql myfirstgosql/main

ENTRYPOINT ["./mygosql"]