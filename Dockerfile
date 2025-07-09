## Build stage
FROM golang:1.23-alpine AS build-env
WORKDIR /app
ADD ./main.go /app/main.go
ADD ./core /app/core
ADD ./go.mod /app/go.mod
ADD ./go.sum /app/go.sum

RUN apk add --update --no-cache git
RUN go mod download
RUN go build -o server

## Creating potential production image
FROM alpine
RUN apk update && apk add bash ca-certificates ffmpeg && rm -rf /var/cache/apk/*
WORKDIR /app
COPY --from=build-env /app/ /app/
COPY ./build/rtsp-stream.yml /app/rtsp-stream.yml
ENTRYPOINT [ "/app/server" ]
