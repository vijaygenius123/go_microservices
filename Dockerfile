FROM golang:1.14.2 AS build-env

ENV GOFLAGS="-mod=vendor"

ADD . /dockerenv
WORKDIR /dockerenv

RUN go build -o /server

FROM debian:buster

EXPOSE 8080

WORKDIR /

COPY --from=build-env /server /

CMD ["/server"]

