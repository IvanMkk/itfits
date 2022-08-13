FROM golang:alpine AS build-env

COPY . .
RUN GOPATH=/ GO111MODULE=on go mod verify
RUN GOPATH=/ GO111MODULE=on GOOS=linux GOARCH=amd64 \
    go build -o bin/itfits

FROM alpine:latest
RUN apk add --update --no-cache ca-certificates
COPY --from=build-env bin/itfits /itfits
ENTRYPOINT ["/itfits"]
