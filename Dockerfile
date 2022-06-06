FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM golang:1.16.6-alpine3.14 AS builder

ARG BUILD_TAG=mongo

WORKDIR /build
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -tags $BUILD_TAG -o structor  -ldflags='-s -w'  -installsuffix cgo ./cmd/structor/.

FROM scratch
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

WORKDIR /structor
COPY --from=builder ./build/structor ./cmd/structor