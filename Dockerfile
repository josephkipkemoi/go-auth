# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /go-auth/

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go mod init go-auth/go-auth-api

RUN go build -o ./go-auth

EXPOSE 8080

CMD ["go-auth"]