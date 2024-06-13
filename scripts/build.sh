#!/usr/bin/env bash

set -e

if [[ "$#" -ne 4 ]]; then
	echo "Usage: $0 <BIN> <VERSION> <OS> <ARCH>"
	exit 1
fi

BIN="$1"
VERSION="$2"
OS="$3"
ARCH="$4"

IMAGE_NAME="$BIN-$OS-$ARCH-image"
CONTAINER_NAME="$BIN-$OS-$ARCH"
OUTPUT_BIN="bin/$BIN"

[[ "$OS" == "darwin" ]] && OUTPUT_BIN="$OUTPUT_BIN-macos"
[[ "$ARCH" == "arm64" ]] && OUTPUT_BIN="$OUTPUT_BIN-arm64"
[[ "$OS" == "windows" ]] && OUTPUT_BIN="$OUTPUT_BIN.exe"

mkdir -p bin

docker build -t "$IMAGE_NAME" \
             --build-arg "OS=$OS" \
             --build-arg "ARCH=$ARCH" \
             --build-arg "BIN=$BIN" \
             --build-arg "VERSION=$VERSION" \
             --no-cache .

docker run --name "$CONTAINER_NAME" "$IMAGE_NAME"
docker cp "$CONTAINER_NAME:/app/$BIN" "$OUTPUT_BIN"
docker rm "$CONTAINER_NAME"
docker image rm "$IMAGE_NAME"
