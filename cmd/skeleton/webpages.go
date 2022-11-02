package main

import (
	"fmt"

	. "github.com/Michael-F-Ellis/gohtx" // dot import makes sense here
)

// indexPage creates the index page as a gohtx HTMLTree. The created
// page will contain key as an hx-val to be used as a session key.
func indexPage(key string) (page *HtmlTree) {
	// We use the Null pseudo-tag here to place the doctype
	// outside the content of the html tag.
	page = Null(
		"<!DOCTYPE html>",
		Html(``,
			Head(``,
				CustomHeadContent(true, true, true),
				Title(``, `Skeleton App`),
			),
			indexBody(key),
		),
	)
	return
}

// indexBody returns the body element of the index.html page
func indexBody(key string) (body *HtmlTree) {
	sectionAttrs := fmt.Sprintf(`class=section hx-vals='{"key": "%s"}'`, key)
	body = Body(``,
		Section(sectionAttrs,
			H1(`class=title`, "Gohtx App Skeleton"),
			P(`class="subtitle is-info"`,
				`with <b>HTMX</b> hypertext and <b>Bulma</b> CSS`,
			),
			Div(`id=target class="container"`,
				Div(`class="block"`, "I've never been updated!"),
				updaterButton(),
			),
		),
	)
	return
}

// updaterButton returns a div containing a button with the htmx attributes
// needed to replace the content of the button's container. The button also
// has HyperScript that toggles the text color of the page title on each click.
func updaterButton() (div *HtmlTree) {
	div = Div(`class="block"`,
		Button(`class="button is-primary is-medium" 
		hx-get="/update" hx-target="#target"
		script="on click toggle .has-text-primary on .title"
		`, "Click Me!"),
	)
	return
}

// updateResponse response returns an html fragment containing a
// message string about the number of updates in the current session
// and an updater button.
func updateResponse(updates uint64) (content *HtmlTree) {
	var msg string
	switch updates {
	case 1:
		msg = `I've been updated once!`
	default:
		msg = fmt.Sprintf(`I've been updated %d times!`, updates)
	}
	content = Null(
		Div(`class="block"`, msg),
		updaterButton(),
	)
	return
}
