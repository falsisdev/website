#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TAILWIND_BIN="$ROOT_DIR/.bin/tailwindcss"

"$ROOT_DIR/build.sh"

cleanup() {
	kill "$TAILWIND_PID" "$SERVER_PID" 2>/dev/null || true
	wait "$TAILWIND_PID" "$SERVER_PID" 2>/dev/null || true
}
trap cleanup EXIT INT TERM

cd "$ROOT_DIR"
"$TAILWIND_BIN" \
	-i web/static/css/input.css \
	-o web/static/css/output.css \
	--watch &
TAILWIND_PID=$!

go run ./cmd/server &
SERVER_PID=$!

wait "$SERVER_PID"
