#!/bin/sh
# Install golangci-lint into $GOPATH/bin and add it to PATH for later CI steps.
# The version must match the `version:` in .custom-gcl.yml so `golangci-lint
# custom` builds a compatible binary. Override by passing a tag as $1.
#
# Built with `go install` (verified by the Go checksum database) rather than the
# upstream install.sh, whose checksum lookup greps the archive name and also
# matches the `.tar.gz.sbom.json` entry, failing verification since v2.15.0.
set -eu

version="${1:-v2.15.0}"
bindir="$(go env GOPATH)/bin"

go install "github.com/golangci/golangci-lint/v2/cmd/golangci-lint@${version}"
"$bindir/golangci-lint" --version

if [ -n "${GITHUB_PATH:-}" ]; then
    echo "$bindir" >> "$GITHUB_PATH"
fi
