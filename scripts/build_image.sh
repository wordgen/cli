#!/usr/bin/env bash

set -e

if [[ "$#" -ne 2 ]]; then
	echo "Usage: $0 <OS> <ARCH>"
	exit 1
fi

OS="$1"
ARCH="$2"
IMAGE_NAME="wordgen-$OS-$ARCH-image"
CONTAINER_NAME="wordgen-$OS-$ARCH"

mkdir -p ./bin

docker build -t "$IMAGE_NAME" \
             --build-arg "OS=$OS" \
             --build-arg "ARCH=$ARCH" \
             --no-cache .

docker run --name "$CONTAINER_NAME" "$IMAGE_NAME"
docker cp "$CONTAINER_NAME:/out/." ./bin
docker rm -f "$CONTAINER_NAME"
