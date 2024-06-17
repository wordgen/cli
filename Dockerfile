FROM golang:1.22.4

ARG OS
ARG ARCH
ARG VERSION

WORKDIR /app

COPY go.mod go.sum ./
COPY main.go flags.go util.go ./
COPY LICENSE ./

RUN go mod tidy

RUN mkdir -p /out

RUN OUTPUT_BIN="wordgen"; \
    if [ "$OS" = "windows" ]; then \
        OUTPUT_BIN="$OUTPUT_BIN.exe"; \
    fi; \
    GOOS=$OS GOARCH=$ARCH go build -o "$OUTPUT_BIN" -trimpath -ldflags="-s -w -X main.version=$VERSION -buildid=" . && \
    ARCHIVE_NAME="wordgen-$VERSION-$OS-$ARCH"; \
    if [ "$OS" = "windows" ]; then \
        apt-get update && apt-get install -y zip; \
        zip -9 "/out/$ARCHIVE_NAME.zip" "LICENSE" "$OUTPUT_BIN"; \
    else \
        tar -cf "/out/$ARCHIVE_NAME.tar.gz" -I "gzip -9" "LICENSE" "$OUTPUT_BIN"; \
    fi

RUN cd /out && for file in *; do sha512sum "$file" > "$file.sha512"; done
