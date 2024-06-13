#!/usr/bin/env bash

set -e

PKGBUILDS=("$HOME/git/PKGBUILDS/wordgen/PKGBUILD"
           "$HOME/git/PKGBUILDS/wordgen-bin/PKGBUILD")
VERSION="$(grep -E '^VERSION[[:space:]]*:=[[:space:]]*' Makefile | awk '{print $3}')"
REPO_URL="https://github.com/wordgen/cli"
TEMP_DIR="$(mktemp -d)"

trap 'rm -rf "$TEMP_DIR"' EXIT INT TERM

for pkgbuild in "${PKGBUILDS[@]}"; do
	cp "$pkgbuild" "$TEMP_DIR"
	cd "$TEMP_DIR"

	sed -i "s/^pkgver=.*/pkgver=${VERSION#v}/" PKGBUILD
	sed -i "s/^pkgrel=.*/pkgrel=1/" PKGBUILD

	if ! updpkgsums; then
		echo "ERROR: updpkgsums failed for $pkgbuild"
		exit 1
	fi

	find . -maxdepth 1 -mindepth 1 ! -name 'PKGBUILD' -exec rm -rf {} +

	if ! makepkg --clean; then
		echo "ERROR: makepkg --clean failed for $pkgbuild"
		exit 1
	fi

	mv PKGBUILD "$pkgbuild"
	cd "$(dirname "$pkgbuild")"

	makepkg --printsrcinfo > .SRCINFO

	git add PKGBUILD .SRCINFO
	git commit -m "$VERSION" -m "- $REPO_URL/releases/tag/$VERSION"
	git push
done
