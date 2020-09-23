FROM golang:alpine

ARG appName
ARG appVersion

ENV app="$appName-$appVersion"
RUN apk update && apk add git

WORKDIR /app

COPY . .
RUN go mod tidy
RUN go build -o $app myfirstgosql/main

ENTRYPOINT ./$app c
