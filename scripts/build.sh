#!/usr/bin/env bash

set -e

if [[ "$#" -ne 2 ]]; then
	echo "Usage: $0 <OS> <ARCH>"
	exit 1
fi

OS="$1"
ARCH="$2"

VERSION="$(grep -E '^VERSION[[:space:]]*:=[[:space:]]*' Makefile | awk '{print $3}')"
IMAGE_NAME="wordgen-$OS-$ARCH-image"
CONTAINER_NAME="wordgen-$OS-$ARCH"
BINARY_NAME="wordgen-$OS-$ARCH"

[[ "$OS" == "windows" ]] && BINARY_NAME="$BINARY_NAME.exe"

mkdir -p bin

docker build -t "$IMAGE_NAME" \
             --build-arg "OS=$OS" \
             --build-arg "ARCH=$ARCH" \
             --build-arg "VERSION=$VERSION" \
             --no-cache .

docker run --name "$CONTAINER_NAME" "$IMAGE_NAME"
docker cp "$CONTAINER_NAME:/app/$BINARY_NAME" "bin"
docker rm "$CONTAINER_NAME"
docker image rm "$IMAGE_NAME"
