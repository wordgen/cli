FROM golang:1.22.4

ARG OS
ARG ARCH
ARG BIN
ARG VERSION

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go ./

RUN go mod tidy

RUN env GOOS=$OS GOARCH=$ARCH go build -o $BIN -trimpath -ldflags="-s -w -X main.version=$VERSION -buildid=" .
