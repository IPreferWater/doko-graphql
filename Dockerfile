ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /doko-graphql
WORKDIR /doko-graphql

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

EXPOSE 8080

#COPY config.json .
ENTRYPOINT ["./app"]