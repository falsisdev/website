package templates

import "embed"

// FS contains the HTML templates used by the application.
//
//go:embed *.html layouts/*.html components/*.html
var FS embed.FS
