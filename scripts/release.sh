#!/usr/bin/env bash

set -e

TAG="$(grep -E '^VERSION[[:space:]]*:=[[:space:]]*' Makefile | awk '{print $3}')"

git tag -s "$TAG" -m "Release $TAG"
git push origin "$TAG"

cd bin

b2sum -- * > b2sums.txt
gpg --local-user FBE12A89 --detach-sign --armor b2sums.txt

FILES=(*)

gh release new "$TAG" --verify-tag "${FILES[@]}"
