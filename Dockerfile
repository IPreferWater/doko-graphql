ARG GO_VERSION=1.14

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /doko-graphql
WORKDIR /doko-graphql

COPY go.mod .
COPY go.sum .

# if you want to add your own certs
#ADD certs /etc/ssl/certs/
RUN mkdir -p etc/ssl/certs
RUN openssl genrsa -out /etc/ssl/certs/server.key 2048
RUN openssl req -new -x509 -key /etc/ssl/certs/server.key \
    -subj "/C=FR/ST=Nord/L=Lille/O=Dis/CN=www.example.com" \
     -out /etc/ssl/certs/server.pem -days 365

RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

EXPOSE 8000

#COPY config.json .
ENTRYPOINT ["./app"]