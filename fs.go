package gohtx

import (
	"embed"
)

//go:embed htmx.min.js bulma/*
var staticFiles embed.FS

// GohtxFS returns the embedded filesystem containing htmx.min.js and bulma.css.
// To serve them, do something like:
// fs := GohtxFS()
// buf, err := fs.ReadFile("htmx.min.js")
func GohtxFS() embed.FS {
	return staticFiles
}

// DefaultHeadContent returns Meta, Link and Script elementss you would
// typically include in the Head element of a page that uses HTMX and Bulma.
func DefaultHeadContent() (h *HtmlTree) {
	h = Null(
		Meta(`charset="utf-8"`),
		Meta(`name="viewport" content="width=device-width, initial-scale=1"`),
		Script(`src="gohtx/htmx.min.js"`),
		Link(`rel="stylesheet" href="gohtx/bulma/css/bulma.min.css" type="text/css"`),
		// Link(`rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bulma@0.9.4/css/bulma.min.css"`),
	)
	return
}
