FROM golang:1.22.4

ARG OS
ARG ARCH

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go ./
COPY LICENSE ./
COPY scripts/build_pack.sh ./

RUN sh build_pack.sh "$OS" "$ARCH"
