#!/usr/bin/env bash

set -euo pipefail

REPO_DIR="$(git rev-parse --show-toplevel)"
VERSION_FILE="$REPO_DIR/VERSION"

cd "$REPO_DIR"

if [[ "$(git branch --show-current)" != main ]]; then
  echo >&2 'ERROR: releases must be created from main'
  exit 1
fi

if [[ -n "$(git status --porcelain)" ]]; then
  echo >&2 'ERROR: working tree is not clean'
  exit 1
fi

CURRENT_TAG="$(git describe --tags --abbrev=0 --match 'v[0-9]*')"
CURRENT_VERSION="$(<"$VERSION_FILE")"

echo "Current tag:       $CURRENT_TAG"
echo "Current version:   $CURRENT_VERSION"

if [[ "$CURRENT_TAG" != "$CURRENT_VERSION" ]]; then
  echo >&2 'ERROR: VERSION file does not match the latest release tag'
  exit 1
fi

read -rp 'Enter new version: ' NEW_VERSION

if [[ ! "$NEW_VERSION" =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo >&2 'ERROR: version must use the vX.Y.Z format'
  exit 1
fi

if [[ "$NEW_VERSION" == "$CURRENT_VERSION" ]]; then
  echo >&2 'ERROR: version has not changed'
  exit 1
fi

if git rev-parse --quiet --verify "refs/tags/$NEW_VERSION" >/dev/null; then
  echo >&2 "ERROR: tag already exists: $NEW_VERSION"
  exit 1
fi

printf '%s\n' "$NEW_VERSION" >"$VERSION_FILE"

git diff -- "$VERSION_FILE"

read -rp "Commit, tag, and push release $NEW_VERSION? [y/N] " CONFIRM

if [[ "$CONFIRM" != [yY] ]]; then
  git restore -- "$VERSION_FILE"
  echo 'Release cancelled.'
  exit 0
fi

git add -- "$VERSION_FILE"
git commit -m "chore(release): bump version to $NEW_VERSION"
git tag "$NEW_VERSION" -m "Release $NEW_VERSION"
git push --atomic origin main "$NEW_VERSION"
