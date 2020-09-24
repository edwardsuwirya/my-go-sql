FROM golang:alpine

ARG appName
ARG appVersion
ARG port

ENV APP="$appName-$appVersion"

RUN apk update && apk add git

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o $APP myfirstgosql/main

EXPOSE $port

ENTRYPOINT ./$APP
