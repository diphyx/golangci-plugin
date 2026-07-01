#!/bin/sh
# Install golangci-lint into $GOPATH/bin and add it to PATH for later CI steps.
# The version must match the `version:` in .custom-gcl.yml so `golangci-lint
# custom` builds a compatible binary. Override by passing a tag as $1.
set -eu

version="${1:-v2.10.1}"
bindir="$(go env GOPATH)/bin"

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
    | sh -s -- -b "$bindir" "$version"

echo "$bindir" >> "$GITHUB_PATH"
