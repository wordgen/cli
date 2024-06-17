#!/usr/bin/env bash

set -e

TAG="$(grep -E '^VERSION[[:space:]]*:=[[:space:]]*' ./Makefile | awk '{print $3}')"

git tag -s "$TAG" -m "Release $TAG"
git push origin "$TAG"

cd ./bin

for file in *.sha512; do
	gpg --local-user FBE12A89 --detach-sign --armor "$file"
done

FILES=(*)

gh release new "$TAG" --verify-tag "${FILES[@]}"
