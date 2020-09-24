FROM golang:alpine

ARG appName
ARG appVersion
ARG port

ENV app="$appName-$appVersion"
RUN apk update && apk add git

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o $app myfirstgosql/main

EXPOSE $port

ENTRYPOINT ./$app h
