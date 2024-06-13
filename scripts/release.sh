#!/usr/bin/env bash

set -e

TAG="$(grep -E '^VERSION[[:space:]]*:=[[:space:]]*' Makefile | awk '{print $3}')"

git tag -s "$TAG" -m "Release $TAG"
git push origin "$TAG"

cd bin

BINS=(*)

for bin in "${BINS[@]}"; do
	sha512sum "$bin" > "$bin.sha512"
	gpg --local-user FBE12A89 --detach-sign --armor "$bin.sha512"
done

cp ../LICENSE .

FILES=(*)

gh release new "$TAG" --verify-tag "${FILES[@]}"
