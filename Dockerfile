FROM golang:1.19-alpine AS build

WORKDIR /pizda_bot

COPY ./app ./
COPY ./.env ./.env

RUN go mod download

RUN go build -o /pbot

## Deploy
FROM alpine:3.17

WORKDIR /

COPY --from=build /pbot /pbot
COPY ./.env /.env

RUN curl -fsSL \
        https://raw.githubusercontent.com/pressly/goose/master/install.sh |\
        GOOSE_INSTALL=/usr sh -s v3.5.0

ENTRYPOINT ["/pbot"]