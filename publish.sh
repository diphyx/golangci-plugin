#!/usr/bin/env bash
set -euo pipefail

# Pick the next version, optionally push, and create a GitHub release whose
# tag publishes the module to the Go proxy.
# Go modules are versioned by git tags (vX.Y.Z); there is no version file to
# edit — pushing the tag is the publish.
# Usage: ./publish.sh

# ─── Validate ───

if [[ ! -f go.mod ]]; then
    echo "Error: go.mod not found (run from repo root)"
    exit 1
fi

# ─── Parse current version (latest vX.Y.Z tag) ───

git fetch origin --tags >/dev/null 2>&1 || true

CURRENT=$(git tag -l 'v*' --sort=-v:refname | head -1)
CURRENT="${CURRENT:-v0.0.0}"
IFS='.' read -r MAJOR MINOR PATCH <<< "${CURRENT#v}"
MAJOR="${MAJOR:-0}"; MINOR="${MINOR:-0}"; PATCH="${PATCH:-0}"

# ─── Prompt ───

echo "Current version: ${CURRENT}"
echo ""
echo "  0) skip    → ${CURRENT}"
echo "  1) patch   → v${MAJOR}.${MINOR}.$((PATCH + 1))"
echo "  2) minor   → v${MAJOR}.$((MINOR + 1)).0"
echo "  3) major   → v$((MAJOR + 1)).0.0"
echo ""
read -rp "Select bump type [0-3]: " choice

case "${choice:-0}" in
    0) ;;
    1) PATCH=$((PATCH + 1)) ;;
    2) MINOR=$((MINOR + 1)); PATCH=0 ;;
    3) MAJOR=$((MAJOR + 1)); MINOR=0; PATCH=0 ;;
    *) echo "Error: invalid choice"; exit 1 ;;
esac

VERSION="v${MAJOR}.${MINOR}.${PATCH}"

if [[ "$VERSION" != "$CURRENT" ]]; then
    echo ""
    echo "==> ${CURRENT} → ${VERSION}"
fi

# ─── Action ───

BRANCH=$(git rev-parse --abbrev-ref HEAD)

echo ""
echo "  0) skip"
echo "  1) push     → push to origin/${BRANCH}"
echo "  2) release  → push and create GitHub release ${VERSION} (publishes module)"
echo ""
read -rp "Select action [0-2]: " action

action="${action:-0}"

if [[ "$action" == "0" ]]; then
    exit 0
fi

git push origin "$BRANCH"
echo ""
echo "==> Pushed ${BRANCH}"

if [[ "$action" == "2" ]]; then
    if [[ "$VERSION" == "$CURRENT" ]]; then
        echo ""
        echo "Error: select a version bump (1-3) to create a release"
        exit 1
    fi

    if [[ "$BRANCH" != "main" ]]; then
        echo ""
        echo "Warning: releases must be created from main (current: ${BRANCH})"
        read -rp "Merge ${BRANCH} into main and continue? [y/N]: " confirm
        case "$confirm" in
            y|Y) ;;
            *) exit 1 ;;
        esac

        SOURCE="$BRANCH"
        git checkout main
        git pull --ff-only origin main
        git merge --no-ff "$SOURCE" -m "Merge branch '${SOURCE}'"
        git push origin main
        BRANCH="main"

        echo ""
        echo "==> Merged ${SOURCE} into main"
    fi

    # Verify the module builds and tests pass before tagging a release
    go build ./...
    go test ./...

    # Detect existing release at this tag — retag flow
    EXISTING_NOTES=""
    if gh release view "$VERSION" >/dev/null 2>&1; then
        echo ""
        echo "Warning: release ${VERSION} already exists"
        read -rp "Delete existing release+tag and re-create at current ${BRANCH}? [y/N]: " confirm
        case "$confirm" in
            y|Y) ;;
            *) exit 1 ;;
        esac

        # Preserve existing notes; fall back to auto-gen if there were none
        EXISTING_NOTES=$(gh release view "$VERSION" --json body --jq .body 2>/dev/null || true)

        gh release delete "$VERSION" --yes --cleanup-tag
        git tag -d "$VERSION" 2>/dev/null || true
        git fetch origin --prune --prune-tags >/dev/null 2>&1 || true

        echo ""
        echo "==> Deleted existing release ${VERSION} and tag"
    fi

    if [[ -n "$EXISTING_NOTES" ]]; then
        gh release create "$VERSION" --title "$VERSION" --notes "$EXISTING_NOTES" --target "$BRANCH"
    else
        gh release create "$VERSION" --title "$VERSION" --generate-notes --target "$BRANCH"
    fi
    echo ""
    echo "==> Created release ${VERSION} (module available via the Go proxy)"
fi
