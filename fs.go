package gohtx

import "embed"

//go:embed htmx.min.js bulma.css
var staticFiles embed.FS

// GohtxFS returns the embedded filesystem containing htmx.min.js and bulma.css.
// To serve them, do something like:
// fs := GohtxFS()
// buf, err := fs.ReadFile("htmx.min.js")
func GohtxFS() embed.FS {
	return staticFiles
}
