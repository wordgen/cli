#!/usr/bin/env bash

set -e

PKGBUILD="$HOME/git/PKGBUILDS/wordgen/PKGBUILD"
REPO_URL="https://github.com/wordgen/cli"
VERSION="$(grep -E '^VERSION[[:space:]]*:=[[:space:]]*' Makefile | awk '{print $3}')"
TEMP_DIR="$(mktemp -d)"

trap 'rm -rf $TEMP_DIR' EXIT INT TERM

cp "$PKGBUILD" "$TEMP_DIR"
cd "$TEMP_DIR"

sed -i "s/^pkgver=.*/pkgver=${VERSION#v}/" PKGBUILD
sed -i "s/^pkgrel=.*/pkgrel=1/" PKGBUILD

if ! updpkgsums; then
	echo "ERROR: updpkgsums failed"
	exit 1
fi

rm -f ./*.tar.gz

if ! makepkg --clean; then
	echo "ERROR: makepkg --clean failed"
	exit 1
fi

mv PKGBUILD "$PKGBUILD"
cd "$(dirname "$PKGBUILD")"

makepkg --printsrcinfo > .SRCINFO

git add PKGBUILD .SRCINFO
git commit -m "$VERSION" -m "- $REPO_URL/releases/tag/$VERSION"
git push
