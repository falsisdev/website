#!/bin/bash
curl -sL https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-linux-x64 -o tw
chmod +x tw
./tw -i web/static/css/input.css -o web/static/css/output.css -m