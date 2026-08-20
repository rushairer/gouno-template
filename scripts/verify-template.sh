#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

export MODULE_PATH="example.com/gouno-template-smoke"
export PROJECT_NAME="gouno_template_smoke"

"$root/scripts/render-template.sh" "$tmp/project" "$MODULE_PATH" "$PROJECT_NAME"

cd "$tmp/project"
if [ -d "$root/../gouno" ]; then
  go mod edit -replace github.com/rushairer/gouno="$root/../gouno"
fi
go mod tidy
go mod download all
go test ./...
