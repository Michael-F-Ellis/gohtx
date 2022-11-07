package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strings"

	"github.com/Michael-F-Ellis/gohtx"
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

	// Gohtx embedded assets handler
	gohtx.AddGohtxAssetHandler()

	// The update handler
	http.Handle("/update", http.HandlerFunc(updateHndlr))

	// The input handlers
	http.Handle("/input", http.HandlerFunc(inputHndlr))
	http.Handle("/unwrapped", http.HandlerFunc(unwrappedInputHndlr))

	// Fragment request handler
	http.Handle("/fragment", http.HandlerFunc(fragmentHndlr))

	// Static file request handler
	http.Handle("/static/", http.FileServer(http.Dir("static")))

	log.Printf("about to serve on %s\n", HostPort)
	switch {
	case strings.HasSuffix(HostPort, ":443"):
		// Live installations for customers serve over HTTPS on port 443 For
		// testing we serve on localhost (default port 8080). The choice of port
		// is determined at startup from a command line argument to the app (see
		// main.go).  Serving securely requires a valid certificate and key. If
		// we can't find paths to both, we log a message and exit.  Obtaining
		// and/or renewing a valid certificate is left up to the service
		// definition that launches the app.

		if err := http.ListenAndServeTLS(HostPort, CertPath, CertKeyPath, nil); err != nil {
			// The usual cause for failure is an expired certificate but other
			// causes are possible.  The value in err will contain an
			// explanation. Log messages are viewable with the journalctl
			// utility.
			log.Fatalf("Could not listen on port %s : %v", HostPort, err)
		}
	default:
		// Serves http (for local testing)
		if err := http.ListenAndServe(HostPort, nil); err != nil {
			// The usual cause for failure in non-secure service is another
			// program or another instance of MARS with a lock on the desired
			// port
			log.Fatalf("Could not listen on port %s : %v", HostPort, err)
		}
	}
}
