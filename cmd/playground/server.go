package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"

	"github.com/Michael-F-Ellis/gohtx"
	"github.com/bitfield/script"
)

// Sessions provides some rudimentary session tracking. You'll
// want something more sophisticated in a real application.
var Sessions = make(map[string]uint64) // key:updateCounter

// newSession initializes a new session. It creates and returns a session key
// string.
func newSession() (key string) {
	key = fmt.Sprintf("%x", rand.Uint64())
	Sessions[key] = 0
	return
}

// Serve serves web pages.
func Serve() {
	// Declare HandlerFuncs

	// The index page handler
	http.Handle("/", http.HandlerFunc(indexHndlr))
	// Embedded assets handler
	// http.Handle("/gohtx/", http.HandlerFunc(gohtxAssetHndlr))
	gohtx.AddGohtxAssetHandler()
	// The update handler
	http.Handle("/update", http.HandlerFunc(updateHndlr))
	// The input handlers
	http.Handle("/input", http.HandlerFunc(inputHndlr))
	http.Handle("/unwrapped", http.HandlerFunc(unwrappedInputHndlr))
	// Fragment request handler
	http.Handle("/fragment", http.HandlerFunc(fragmentHndlr))

	// Serve the static files
	http.Handle("/static/", http.FileServer(http.Dir("static")))

	// Serve the static files

	// Live installations for customers serve over HTTPS on port 443
	// For testing we serve on localhost (default port 8080). The choice of
	// port is determined at startup from a command line argument to the app (see main.go).
	// Serving securely requires a valid certificate and key. If we can't find paths to
	// both, we log a message and exit.  Obtaining and/or renewing a valid certificate is left up to
	// the service definition that launches the app.
	log.Printf("about to serve on %s\n", HostPort)
	switch {
	case strings.HasSuffix(HostPort, ":443"):
		if err := http.ListenAndServeTLS(HostPort, CertPath, CertKeyPath, nil); err != nil {
			// The usual cause for failure is an expired certificate but other
			// causes are possible.  The value in err will contain an
			// explanation. Log messages are viewable with the journalctl
			// utility.
			log.Fatalf("Could not listen on port %s : %v", HostPort, err)
		}
	default:
		if err := http.ListenAndServe(HostPort, nil); err != nil {
			// The usual cause for failure in non-secure service is another
			// program or another instance of MARS with a lock on the desired
			// port
			log.Fatalf("Could not listen on port %s : %v", HostPort, err)
		}
	}
}

// indexHndlr generates and returns the index page.
func indexHndlr(w http.ResponseWriter, r *http.Request) {
	var buf bytes.Buffer
	// For this skeleton, we start a new session
	// when the index page is loaded or reloaded.
	err := gohtx.Render(indexPage(newSession()), &buf, 0)
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

	err := gohtx.Render(updateResponse(count), &buf, 0)
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
	code := r.FormValue("code")
	htm := eval(code, true)
	_, _ = w.Write([]byte(htm))
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
	htm := eval(code, false) // don't insert the user's code into the template.
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
func eval(input string, wrap bool) (htm string) {
	// Insert user input into the template
	var code string
	if wrap {
		code = fmt.Sprintf(template, input)
	} else {
		code = input
	}
	// Get a temporary file to hold the code.
	tmpfile, err := ioutil.TempFile("temp", "*.go")
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
	}
	return
}

var template = `
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

// getFragments returns a map of the files in the fragments FS.
// The file contents are keyed first line of each fragment file.
func getFragments() (m map[string]string, err error) {
	m = make(map[string]string)
	files, err := fragments.ReadDir("fragments")
	if err != nil {
		return
	}
	for _, f := range files {
		buf, e := os.ReadFile("fragments/" + f.Name())
		err = e
		if err != nil {
			return
		}
		code := strings.TrimSpace(string(buf))
		lines := strings.Split(code, "\n")
		name := strings.TrimSpace(strings.ReplaceAll(lines[0], "/", ""))
		m[name] = code
	}
	return
}
