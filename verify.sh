#!/bin/sh
# End-to-end test: run the REAL linter (a custom golangci-lint with the diphyx
# plugin built in) over the fixture packages under example/, and turn the result
# into a deterministic pass/fail check.
#
# golangci-lint exits non-zero when it finds issues, but that alone can't tell us
# a *specific* rule regressed — a rule could silently stop firing and the run
# would still "fail" on the others. So this script asserts both directions:
#
#   1. valid    — example/valid/**  must produce ZERO issues (no false positives).
#   2. invalid  — example/invalid/** must, for EVERY diphyx analyzer:
#                   a. produce no typecheck errors (every fixture must compile),
#                   b. fire that analyzer at least once, and
#                   c. leave no invalid/<rule> package silent (so a fixture that
#                      stops firing can't hide behind a sibling covering the same
#                      analyzer name).
#
# The expected analyzer set is derived from the invalid/ subdirectories, so
# adding a rule + its invalid/<rule> fixture is all it takes to extend coverage.
#
# Run with `make verify` (or `sh verify.sh`). Exits non-zero on failure.

set -u

repo_root=$(CDPATH= cd -- "$(dirname -- "$0")" && pwd)
example_dir="$repo_root/example"
gcl="$repo_root/custom-gcl"

status=0

# --- 0. build the custom golangci-lint (with the diphyx plugin) if missing ---
if [ ! -x "$gcl" ]; then
    echo "• building custom-gcl (golangci-lint custom)…"
    if ! (cd "$repo_root" && golangci-lint custom >/tmp/diphyx-custom-build.log 2>&1); then
        echo "✗ failed to build custom-gcl:"
        cat /tmp/diphyx-custom-build.log
        exit 1
    fi
fi

valid_json=$(mktemp)
invalid_json=$(mktemp)

# --- 1. valid: nothing may warn ---
(cd "$example_dir" && "$gcl" run \
    --output.json.path="$valid_json" --output.text.path=/dev/null \
    ./valid/... >/dev/null 2>/dev/null)

valid_count=$(grep -o '"FromLinter"' "$valid_json" | wc -l | tr -d ' ')
if [ "$valid_count" = "0" ]; then
    echo "✓ valid: 0 issues across example/valid"
else
    echo "✗ valid fixtures produced $valid_count issue(s):"
    grep -oE '"Text":"[^"]+"' "$valid_json" | sed 's/"Text":"/    /; s/"$//'
    status=1
fi

# --- 2. invalid: every rule must fire, no fixture silent, everything compiles ---
(cd "$example_dir" && "$gcl" run \
    --output.json.path="$invalid_json" --output.text.path=/dev/null \
    ./invalid/... >/dev/null 2>/dev/null)

if grep -q '"FromLinter":"typecheck"' "$invalid_json"; then
    echo "✗ invalid fixtures have compile/typecheck errors:"
    grep -oE '"Text":"[^"]+"' "$invalid_json" | sed 's/"Text":"/    /; s/"$//'
    status=1
fi

# Expected analyzers = the invalid/ subdirectories (each named after its rule).
rules=$(ls -1 "$example_dir/invalid")
total=$(printf '%s\n' "$rules" | grep -c .)

missing=""
silent=""
for rule in $rules; do
    # b. the analyzer fired at least once (messages are prefixed "<rule>: …").
    if ! grep -q "\"Text\":\"$rule:" "$invalid_json"; then
        missing="$missing $rule"
    fi

    # c. the invalid/<rule> package itself produced at least one issue.
    if ! grep -q "\"Filename\":\"invalid/$rule/" "$invalid_json"; then
        silent="$silent $rule"
    fi
done

if [ -n "$missing" ]; then
    echo "✗ diphyx rule(s) that never fired:"
    for rule in $missing; do echo "    $rule"; done
    status=1
fi

if [ -n "$silent" ]; then
    echo "✗ invalid fixture package(s) that produced no issue:"
    for rule in $silent; do echo "    invalid/$rule"; done
    status=1
fi

if [ -z "$missing" ] && [ -z "$silent" ]; then
    echo "✓ invalid: all $total diphyx rules fired, no fixture silent"
fi

rm -f "$valid_json" "$invalid_json"

echo
if [ "$status" -ne 0 ]; then
    echo "harness test FAILED"
    exit 1
fi
echo "harness test passed"
