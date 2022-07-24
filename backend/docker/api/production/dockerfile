FROM golang:1.18

WORKDIR /go/api

COPY go.* ./
RUN go mod download

COPY . ./

RUN go build -o /api