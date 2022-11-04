package gohtx

import (
	"embed"
	"log"
	"net/http"
	"path"
	"strings"
)

const GohtxAssetPath = "/gohtx/" // default route to embedded assets.

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

// AddGohtxAssetHandler is a convenience function that adds an http.Handler to
// process requests for assets in the embedded FS. Use this to avoid having to
// know about the path string.
func AddGohtxAssetHandler() {
	http.Handle(GohtxAssetPath, http.HandlerFunc(assetHandler))
}

// assetHandler is the unexported handler installed by GohtxAssetHandler
func assetHandler(w http.ResponseWriter, r *http.Request) {
	asset := strings.TrimPrefix(r.URL.Path, GohtxAssetPath)
	log.Printf("%s requested", asset)
	fs := GohtxFS()
	buf, err := fs.ReadFile(asset)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Write the Content-Type to the header
	ext := path.Ext(asset)
	switch ext {
	case ".css":
		w.Header().Set("Content-Type", "text/css")
	case ".js":
		w.Header().Set("Content-Type", "text/javascript")
	default:
		w.Header().Set("Content-Type", http.DetectContentType(buf))
	}
	_, err = w.Write(buf)
	if err != nil {
		log.Println(err)
	}
}
