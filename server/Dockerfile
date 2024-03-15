FROM golang:1.20-buster as builder

RUN apt-get update && \
    apt-get -qy full-upgrade && \
    apt-get install -qy curl && \
    apt-get install -qy apt-transport-https ca-certificates gnupg

WORKDIR /src
COPY go.* ./
RUN go mod download
COPY . ./

RUN go build -v -o /bin/server