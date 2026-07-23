#!/bin/bash
curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64
chmod +x tailwindcss-linux-x64
./tailwindcss-linux-x64 -i ./web/static/css/input.css -o ./web/static/css/output.css --minify
go build -o server cmd/server/main.go