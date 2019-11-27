ARG GO_VERSION=1.13

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o ./app ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /api
WORKDIR /api
COPY --from=builder /api .

ENV CONNSTR_DRVR mysql
ENV CONNSTR_UID root
ENV CONNSTR_PASS  #P@ssw0rd
ENV CONNSTR_SERVER 127.0.0.1
ENV CONNSTR_PORT 3366
ENV CONNSTR_DBNAME promo

ENV SERVER_HOST 127.0.0.1
ENV SERVER_PORT 8080
ENV SERVER_MODE debug

EXPOSE 8080

ENTRYPOINT ["./app"]