#!/usr/bin/env bash
set -euo pipefail

root="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
tmp="$(mktemp -d)"
trap 'rm -rf "$tmp"' EXIT

export MODULE_PATH="example.com/gouno-template-smoke"
export PROJECT_NAME="gouno_template_smoke"

"$root/scripts/render-template.sh" "$tmp/project" "$MODULE_PATH" "$PROJECT_NAME"

cd "$tmp/project"
if [ -n "${GOUNO_REPLACE_DIR:-}" ]; then
  test -f "$GOUNO_REPLACE_DIR/go.mod"
  go mod edit -replace github.com/rushairer/gouno="$GOUNO_REPLACE_DIR"
fi
cp go.mod "$tmp/go.mod.before"
cp go.sum "$tmp/go.sum.before"
go mod tidy
cmp -s go.mod "$tmp/go.mod.before"
cmp -s go.sum "$tmp/go.sum.before"
go mod download all
gofmt -w .
test -z "$(gofmt -l .)"
go test -race ./...
go vet ./...
go run golang.org/x/vuln/cmd/govulncheck@v1.6.0 ./...
go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.12.2 run
