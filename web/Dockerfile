FROM golang:1.19-alpine AS build

WORKDIR /web

COPY ./web ./

RUN go mod tidy
RUN go mod download
RUN go build -o /webserver


## Deploy
FROM alpine:3.17

WORKDIR /

COPY --from=build /webserver /webserver
COPY ./.env /.env

ENTRYPOINT ["/webserver"]
#CMD ["tail", "-f", "/dev/null"]