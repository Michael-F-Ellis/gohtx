package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	. "github.com/Michael-F-Ellis/gohtx" // dot import makes sense here
	"github.com/bitfield/script"
)

// indexHndlr generates and returns the index page.
func indexHndlr(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	// For this skeleton, we start a new session
	// when the index page is loaded or reloaded.
	err := Render(indexPage(newSession()), &buf, 0)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, err = w.Write(buf.Bytes())
		if err != nil {
			log.Println(err)
		}
	}
}

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
				// Set textarea heights to automatically resize on input.
				Script(`type=text/javascript`, `
		function AutomateTextareaYScroll() {
		var tx = document.getElementsByTagName("textarea");
        for (let i = 0; i < tx.length; i++) {
        tx[i].setAttribute("style", "height:" + (tx[i].scrollHeight) + "px;overflow-y:hidden;");	
		tx[i].addEventListener("input", OnInput, false);
        }}

        function OnInput() {
        this.style.height = 0;
        this.style.height = (this.scrollHeight) + "px";
        }
		htmx.onLoad(function(elt) {AutomateTextareaYScroll();});
		`),
			),
			indexBody(key),
		),
	)
	return
}

// indexBody returns the body element of the index.html page
func indexBody(key string) (body *HtmlTree) {
	sectionAttrs := fmt.Sprintf(`class=section hx-vals='{"key": "%s"}'`, key)
	var optionNames []string
	for name := range Fragments {
		optionNames = append(optionNames, name)
	}
	defaultExample, ok := Fragments["Notification"]
	if !ok {
		// It's a programming error if the fragment named 'notification' isn't available.
		panic("defaultExample 'notification' not found in Fragments map.")
	}
	body = Body(``,
		Section(sectionAttrs,
			// Title and subtitle
			H1(`class="title has-text-centered"`, "Gohtx Playground"),
			P(`class="subtitle is-info has-text-centered"`,
				`with <b>HTMX</b>, <b>HyperScript</b> and <b>Bulma</b> CSS`),
			// A dropdown select of example code fragments
			Div(`class=container`,
				mkSelect(
					optionNames,
					"/fragment",
					"which",
					"#gocode",
				)),
			Div(`class=container id="forms-and-results"`,
				formsAndResultsContent("", defaultExample, false, true, false),
			),
		),

		// documentation links
		resourceLinks(),
	)
	return
}

// formsAndResultsContent returns updated content for the forms-and-results
// block.  Result is a string that contain either html code or Go code. Which
// language is indicated by isHtml. The evaluation that created result may have
// been unsuccessful, as indicated by isOk.  Successful html results are copied
// into #pgtarget to be rendered by the browser and into #pghtml so the user can
// see the generated code. Unsuccessful html results are just a thin wrapper
// around error messages from the evaluation that created the html. We copy
// those only into #pgtarget.
// Go code results are produced by Ghotify from html entered by the user into
// #pghtml.  These results also may be successfull or unsuccessful as indicated
// by isOk.  Successful results are copied into #pgsource only. Unsuccessful
// results go into #pgtarget only. The nullwrap flag determines whether generated
// Go code will be wrapped with a null tag.
func formsAndResultsContent(src, result string, isHtml, isOk bool, nullwrap bool) (content string) {
	var (
		pgSourceContent, pgHtmlContent, pgTargetContent string
	)
	switch isHtml {
	case true:
		pgSourceContent = src
		pgTargetContent = result
		if isOk {
			pgHtmlContent = result
		} else {
			pgSourceContent = src
		}
	case false:
		pgHtmlContent = src
		if isOk {
			if nullwrap {
				pgSourceContent = "htx = " + result
			} else {
				pgSourceContent = result
			}
			pgTargetContent = ""
		} else {
			pgSourceContent = ""
			pgTargetContent = result
		}

	}
	htree := Null(
		// Gohtx code
		Div(`id="pgsource" class="block"`,
			// A form with textarea for code and a button to submit it.
			Form(`id="pgsrcform" class="form" hx-post="/input" hx-target="#forms-and-results" hx-vals='{"lang":"go"}'`,
				labeledFormField("Go Code", Textarea(`id="gocode" class=textarea name=gocode`, pgSourceContent)),
				unlabeledFormField(Button(`class="button is-primary" type="submit"`, "Evaluate")),
			),
		),
		// where the server response goes
		Div(`id="pgtarget" class="block"`, pgTargetContent),

		// Html code
		Div(`id="pghtml" class="block"`,
			Form(`class="form" hx-post="/input" hx-target="#forms-and-results" hx-vals='{"lang":"html"}'`,
				labeledFormField("HTML", Textarea(`id="htmlcode" class=textarea name=htmlcode`, pgHtmlContent)),
				unlabeledFormField(Button(`class="button is-primary" type="submit"`, "Gohtify")),
			),
		),
	)
	var buf bytes.Buffer
	err := Render(htree, &buf, 0)
	if err != nil {
		content = fmt.Sprintf("Unable to render formcontent: %v", err)
	} else {
		content = buf.String()
	}
	return
}

// resourceLinks returns a div with links to documentation for htmx,
// hyperscript, bulma and gohtx.
func resourceLinks() (div *HtmlTree) {
	div = Div(`class="container"`,
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

func mkSelect(optionNames []string, url, param, target string) (sel *HtmlTree) {
	var options []interface{}
	for _, name := range optionNames {
		attrs := fmt.Sprintf(`value="%s"`, name)
		options = append(options, Option(attrs, name))
	}
	attrs := fmt.Sprintf(`name="%s" hx-get="%s" hx-target="%s"`, param, url, target)
	sel = labeledFormField("Examples", Div(`class="select is-link"`,
		Select(attrs, options...)))
	return
}

// updateHndlr responds to an update request. It verifies that
// the request contains a valid session key before generating
// and rendering the html.
func updateHndlr(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	key := r.URL.Query().Get("key")
	if key == "" {
		log.Println("No key in request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	count, ok := Sessions[key]
	if !ok {
		log.Printf("Invalid key in request: %s", key)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	count++
	Sessions[key] = count

	err := Render(updateResponse(count), &buf, 0)
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		_, err = w.Write(buf.Bytes())
		if err != nil {
			log.Println(err)
		}
	}
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

// inputHndlr gets the user's Go code from the input textarea and tries to evaluate
// it. It uses the eval function which puts the result of the evaluation into the supplied
// buffer.
func inputHndlr(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	lang := r.FormValue("lang")
	var (
		result string
	)
	switch lang {
	case "go":
		// If the go cold evaluates successfully, we put the resulting hmtl into two places.
		// One
		code := r.FormValue("gocode")
		evalResult, ok := eval(code, true)
		result = formsAndResultsContent(code, evalResult, true, ok, false)
	case "html":
		code := r.FormValue("htmlcode")
		ignore := map[string]struct{}{"html": {}, "head": {}, "body": {}}
		err := Gohtify(code, true, ignore, &result)
		if err != nil {
			log.Printf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		result = formsAndResultsContent(code, result, false, true, true)
	default:
		log.Printf("unknown lang:'%v'", lang)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, _ = w.Write([]byte(result))
}

// unwrappedInputHndlr gets the user's Go code from the input textarea and tries to evaluate
// it. It uses the eval function which puts the result of the evaluation into the supplied
// buffer.
func unwrappedInputHndlr(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Printf("%v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	code := r.FormValue("code")
	htm, _ := eval(code, false) // don't insert the user's code into the template.
	_, _ = w.Write([]byte(htm))
}

// fragmentRequestHndlr returns the request code fragment
func fragmentHndlr(w http.ResponseWriter, r *http.Request) {
	which := r.URL.Query().Get("which")
	log.Println(which)
	code, ok := Fragments[which]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(code)
	_, _ = w.Write([]byte(code))
}

// eval is called to evaluate Go code entered in the playground.
func eval(input string, wrap bool) (htm string, ok bool) {
	// Insert user input into the template
	var code string
	if wrap {
		code = fmt.Sprintf(wrapper, input)
	} else {
		code = input
	}
	// Get a temporary file to hold the code.
	tmpfile, err := os.CreateTemp("temp", "*.go")
	if err != nil {
		htm = fmt.Sprintf("<p>%v</p>", err)
		return
	}
	defer os.Remove(tmpfile.Name())

	// Write the code to the temporary file.
	if _, err := tmpfile.Write([]byte(code)); err != nil {
		htm = fmt.Sprintf("<p>%v</p>", err)
		return
	}
	// Run the code
	cmd := fmt.Sprintf("go run %s", tmpfile.Name())
	htm, err = script.Exec(cmd).String()
	if err != nil {
		// Atempt to format the code but don't worry if formatting fails.
		p := script.Exec("go fmt " + tmpfile.Name())
		p.Wait()
		// Return a listing of the run errors and the code.
		htm = fmt.Sprintf(`
		<div class="notification is-warning">
		  <p>Evaluation failed:</p>
		  <pre class="notification is-warning">%v</pre>
		</div>
		<hr><code><pre>%v</pre></code>`, htm, code)

		return
	}
	ok = true
	return
}

// Wrapper contains a small main program text that wraps around code fragments
// submitted for evaluation. The program is designed to be executed by 'go run'.
// It attempts to render the code fragment and print it to stdout. If there is
// an error, an html <p> containing the error message is printed instead.
var wrapper = `
// Wrapper for Gohtx Playground evaluation
package main

import (
    "bytes"
    "fmt"

    . "github.com/Michael-F-Ellis/gohtx"
)

func main() {
    var htx *HtmlTree
	htx = P("",B("", "If you see this message, you forgot to assign a value to 'htx'."))

    /***** You code inserted here. Must assign a *HtmlTree to htx. *****/
    %s
    /***** End of your code. *****/

    var buf bytes.Buffer
    err := Render(htx, &buf, 0)
    if err != nil {
    	buf.Reset()
    	err = Render(P("", "render failed: "+err.Error()), &buf, 0)
    	if err != nil {
    	    // This should never happen ...
    		panic(err)
    	}
    }
    fmt.Println(buf.String())
}`
