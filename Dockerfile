# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY . ./

RUN apk update \
    && apk add ffmpeg

ARG VERSION

RUN go mod download

RUN go build -o /garrus -ldflags "-X `go list ./version`.VersionStr=${VERSION}"

CMD [ "/garrus" ]