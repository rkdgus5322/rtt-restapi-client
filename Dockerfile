# syntax=docker/dockerfile:1

FROM golang:alpine AS builder

RUN mkdir -p /go/src/app

WORKDIR /go/src/app
COPY . .

ENV CGO_ENABLED=0
RUN go get
RUN go mod download
RUN go build -a -o rtt-client

FROM alpine

WORKDIR /app

COPY --from=builder /go/src/app/rtt-client /app/

ENV GO111MODULE="on"
ENV GIN_MODE="release"

EXPOSE 8080

CMD ["./rtt-client"]