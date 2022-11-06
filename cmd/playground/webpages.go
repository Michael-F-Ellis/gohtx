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
				Title(``, `Gohtx Playground`),
			),
			indexBody(key),
		),
	)
	return
}

// labeledFormField wraps a label and control into a bulma form field.
func labeledFormField(label string, control *HtmlTree) (field *HtmlTree) {
	field = Div(`class="field"`,
		Label(`class="label"`, label),
		Div(`class="control"`, control))
	return
}

// unlabeledFormField wraps a control into a bulma form field.
func unlabeledFormField(control *HtmlTree) (field *HtmlTree) {
	field = Div(`class="field"`,
		Div(`class="control"`, control))
	return
}

// indexBody returns the body element of the index.html page
func indexBody(key string) (body *HtmlTree) {
	defaultExample, err := fragments.ReadFile("fragments/notification.txt")
	if err != nil {
		// It's a programming error if notifications.txt isn't available.
		panic(err)
	}
	sectionAttrs := fmt.Sprintf(`class=section hx-vals='{"key": "%s"}'`, key)

	body = Body(``,
		Section(sectionAttrs,
			// Title and subtitle
			H1(`class="title has-text-centered"`, "Gohtx Playground"),
			P(`class="subtitle is-info has-text-centered"`,
				`with <b>HTMX</b>, <b>HyperScript</b> and <b>Bulma</b> CSS`),

			Div(`class="block"`,
				// A form with textarea for code and a button to submit it.
				Div(`id="pgsource" class="block"`,
					Form(`class="form" hx-post="/input" hx-target="#pgtarget"`,
						labeledFormField("Code", Textarea(`class=textarea name=code`, string(defaultExample))),
						unlabeledFormField(Button(`class="button is-primary" type="submit"`, "Evaluate"))),
				),
			),
			// where the server response goes
			Div(`id="pgtarget" class="block"`),
		),

		Div(`class="container"`,
			Div(`class="block"`,
				Hr(``),
				P(``, `Learn more about HTMX, HyperScript, and Bulma at:`),
				Ul(``,
					Li(``, A(`href="https://htmx.org"`, "htmx.org")),
					Li(``, A(`href="https://hyperscript.org"`, "hyperscript.org")),
					Li(``, A(`href="https://bulma.io"`, "bulma.io")),
				),
			),
			Div(`class="block"`,
				P(``, `Learn more about Gohtx at:`),
				Ul(``,
					Li(``, A(`href="https://pkg.go.dev/github.com/Michael-F-Ellis/gohtx"`, "pkg.go.dev")),
					Li(``, A(`href="https://github.com/Michael-F-Ellis/gohtx"`, "github.com")),
				),
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
