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
BINARY_NAME="$BIN-$OS-$ARCH"

[[ "$OS" == "windows" ]] && BINARY_NAME="$BINARY_NAME.exe"

mkdir -p bin

docker build -t "$IMAGE_NAME" \
             --build-arg "OS=$OS" \
             --build-arg "ARCH=$ARCH" \
             --build-arg "BIN=$BIN" \
             --build-arg "VERSION=$VERSION" \
             --no-cache .

docker run --name "$CONTAINER_NAME" "$IMAGE_NAME"
docker cp "$CONTAINER_NAME:/app/$BINARY_NAME" "bin"
docker rm "$CONTAINER_NAME"
docker image rm "$IMAGE_NAME"
