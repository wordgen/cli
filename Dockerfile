FROM golang:1.22.4

ARG OS
ARG ARCH
ARG VERSION

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go flags.go util.go ./

RUN go mod tidy

RUN OUTPUT_BIN="wordgen-$OS-$ARCH"; \
    if [ "$OS" = "windows" ]; then \
        OUTPUT_BIN="$OUTPUT_BIN.exe"; \
    fi; \
    GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUT_BIN" -trimpath -ldflags="-s -w -X main.version=$VERSION -buildid=" .
