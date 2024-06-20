#!/usr/bin/env sh

set -e

OS="$1"
ARCH="$2"
VERSION="$(grep -oP 'version\s*=\s*"\K[^"]+' ./main.go)"
OUTPUT_BIN="wordgen"
ARCHIVE_NAME="wordgen-$VERSION-$OS-$ARCH"

if [ "$OS" = "windows" ]; then
	OUTPUT_BIN="$OUTPUT_BIN.exe"
fi

GOOS="$OS" GOARCH="$ARCH" go build -o "$OUTPUT_BIN" -trimpath -ldflags="-s -w -buildid=" .

mkdir -p "/out/$ARCHIVE_NAME"
mv "$OUTPUT_BIN" "LICENSE" "/out/$ARCHIVE_NAME"
cd /out

if [ "$OS" = "windows" ]; then
	apt-get update && apt-get install -y zip
	zip -9r "$ARCHIVE_NAME.zip" "$ARCHIVE_NAME"
else
	tar -cf "$ARCHIVE_NAME.tar.gz" -I "gzip -9" "$ARCHIVE_NAME"
fi

rm -rf "$ARCHIVE_NAME"

for file in *; do
	sha512sum "$file" > "$file.sha512"
done
