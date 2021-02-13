FROM golang:1.15.8-alpine AS build

WORKDIR /src/
RUN apk update && apk add --no-cache  \
        bash              \
        build-base        \
        coreutils         \
        gcc               \
        git               \
        make              \
        musl-dev          \
        openssl-dev       \
        openssl           \
        libsasl           \
        libgss-dev        \
        rpm               \
        lz4-dev           \
        zlib-dev          \
        ca-certificates   \
        wget

COPY *.go go.* /src/
RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -o /bin/app -a -v -tags musl
EXPOSE 8080

ENTRYPOINT ["/bin/app"]
