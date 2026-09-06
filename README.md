# Personal Website

A small personal website starter built with Go, `net/http`, Go HTML templates, Tailwind CSS, and HTMX.

## Requirements

- Go 1.23 or newer
- Bash
- `curl` for downloading the Tailwind CLI on the first build

## Development

Run the complete development environment with one command from the project root:

```bash
./dev.sh
```

This command:

- downloads the platform-specific Tailwind CLI into `.bin/` when needed;
- watches the templates and rebuilds `web/static/css/output.css`;
- starts the Go server on `http://localhost:8080`;
- stops both processes when you press `Ctrl+C`.

If the script is not executable after cloning, run:

```bash
chmod +x build.sh dev.sh
```

You can also run the pieces separately:

```bash
./build.sh
.bin/tailwindcss -i web/static/css/input.css -o web/static/css/output.css --watch
go run ./cmd/server
```

Tailwind scans the files under `web/templates`. Add utility classes directly to the templates, and the watch process will update the generated CSS.

## Bilingual interface

The language selector in the navbar supports English and Turkish. Changing the select sends an HTMX request and replaces only `#page-shell`, so the browser does not perform a full page reload.

The selected language is stored in a cookie for one year. Add new translations in `handlers/pageData` and use the fields in the templates.

## Project structure

```text
api/                         Vercel serverless function entrypoint
cmd/server/                  Local Go server entrypoint
handlers/                    Shared HTTP handlers and page data
web/templates/               Embedded HTML templates
web/static/css/input.css     Tailwind input file
build.sh                     Production Tailwind build
dev.sh                       Local Tailwind watch and Go server
vercel.json                  Vercel build and routing configuration
```

Templates are embedded into the Go application with `embed.FS`. This is required because Vercel functions do not guarantee that the repository's `web/templates` path exists at runtime.

## Production build

Build the CSS and the Vercel static output locally with:

```bash
./build.sh
```

The generated files are placed in `public/static/` and are ignored by Git.

## Vercel deployment

Use these project settings in Vercel:

- **Framework Preset:** `Other`
- **Build Command:** `./build.sh`
- **Output Directory:** `public`
- **Install Command:** leave empty
- **Root Directory:** the repository root

The `api/index.go` file is deployed as the Go serverless function. Static files are served from `public/static/`.
