#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TAILWIND_BIN="$ROOT_DIR/.bin/tailwindcss"

case "$(uname -s)-$(uname -m)" in
	Darwin-arm64) TAILWIND_TARGET="macos-arm64" ;;
	Darwin-x86_64) TAILWIND_TARGET="macos-x64" ;;
	Linux-x86_64) TAILWIND_TARGET="linux-x64" ;;
	Linux-aarch64) TAILWIND_TARGET="linux-arm64" ;;
	*) echo "Unsupported platform: $(uname -s)-$(uname -m)" >&2; exit 1 ;;
esac

mkdir -p "$(dirname "$TAILWIND_BIN")"

if [[ ! -x "$TAILWIND_BIN" ]]; then
	curl -fsSL "https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$TAILWIND_TARGET" -o "$TAILWIND_BIN"
	chmod +x "$TAILWIND_BIN"
fi

"$TAILWIND_BIN" \
	-i "$ROOT_DIR/web/static/css/input.css" \
	-o "$ROOT_DIR/web/static/css/output.css" \
	--minify