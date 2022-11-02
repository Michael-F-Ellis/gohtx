package gohtx

import (
	"embed"
)

//go:embed htmx.min.js hyperscript.js bulma/*
var staticFiles embed.FS

// GohtxFS returns the embedded filesystem containing htmx.min.js and bulma.css.
// To serve them, do something like:
// fs := GohtxFS()
// buf, err := fs.ReadFile("htmx.min.js")
func GohtxFS() embed.FS {
	return staticFiles
}

// DefaultHeadContent calls CustomHeadContent specifying htmx and bulma
// without hyperscript.
func DefaultHeadContent() (h *HtmlTree) {
	h = CustomHeadContent(true, false, true)
	return
}

// CustomHeadContent returns the same Meta elements you would typically include
// in the <head> element of a responsive web page and allows you to choose
// which, if any, of htmx, hyperscript, and bulma to include.
func CustomHeadContent(htmx, hyperscript, bulma bool) (h *HtmlTree) {
	var options []interface{}
	options = append(options,
		Meta(`charset="utf-8"`),
		Meta(`name="viewport" content="width=device-width, initial-scale=1"`),
	)

	if htmx {
		options = append(options, Script(`src="gohtx/htmx.min.js"`))
	}
	if hyperscript {
		options = append(options, Script(`src="gohtx/hyperscript.js"`))
	}
	if bulma {
		options = append(options,
			Link(`rel="stylesheet" href="gohtx/bulma/css/bulma.min.css" type="text/css"`))
	}

	h = Null(options...)
	return
}
