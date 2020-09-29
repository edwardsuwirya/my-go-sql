FROM golang:alpine as build-env

ARG appName
ARG appVersion
ARG port

ENV APP="$appName-$appVersion"

RUN apk update && apk add git

WORKDIR /src

COPY . .

RUN go mod tidy
RUN go build -o $APP myfirstgosql/main

FROM alpine
WORKDIR /app
COPY --from=build-env /src/$APP /app

EXPOSE $port

ENTRYPOINT ./$APP
