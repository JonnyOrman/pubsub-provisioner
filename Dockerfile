FROM golang:1.18-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

COPY . /pubsub-provisioner

WORKDIR /pubsub-provisioner

RUN go mod download

RUN go build -o ./pubsub-provisioner ./main.go

EXPOSE 8080

ENTRYPOINT ["./pubsub-provisioner"]