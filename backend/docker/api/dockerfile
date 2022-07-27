FROM golang:1.18

WORKDIR /go/api

COPY ./ .

RUN go mod download
RUN go install github.com/cosmtrek/air@latest