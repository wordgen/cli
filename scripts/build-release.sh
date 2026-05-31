#!/usr/bin/env bash

set -euo pipefail

VERSION="${1:?usage: $0 <version>}"

REPO_DIR="$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." && pwd)"
DIST_DIR="$REPO_DIR/dist"
STAGE_DIR="$(mktemp -d)"

trap 'rm -rf "$STAGE_DIR"' EXIT

targets=(
    linux/amd64
    linux/arm64
    windows/amd64
    windows/arm64
    darwin/amd64
    darwin/arm64
)

rm -rf "$DIST_DIR"
mkdir -p "$DIST_DIR"

cd "$REPO_DIR"

for target in "${targets[@]}"; do
    os="${target%/*}"
    arch="${target#*/}"
    name="wordgen-$VERSION-$os-$arch"
    package_dir="$STAGE_DIR/$name"
    binary="$package_dir/wordgen"

    if [[ "$os" == windows ]]; then
        binary+='.exe'
    fi

    mkdir -p "$package_dir"

    CGO_ENABLED=0 GOOS="$os" GOARCH="$arch" \
        go build \
        -trimpath \
        -ldflags='-s -w -buildid=' \
        -o "$binary" \
        ./cmd/wordgen

    cp LICENSE "$package_dir/"

    if [[ "$os" == windows ]]; then
        cd "$STAGE_DIR"
        zip -qr "$DIST_DIR/$name.zip" "$name"
        cd "$REPO_DIR"
    else
        tar -C "$STAGE_DIR" -czf "$DIST_DIR/$name.tar.gz" "$name"
    fi

    rm -rf "$package_dir"
done

cd "$DIST_DIR"

sha512sum wordgen-* >SHA512SUMS
sha512sum -c SHA512SUMS
