package main

import (
	"bytes"
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
	// Embedded assets handler
	// http.Handle("/gohtx/", http.HandlerFunc(gohtxAssetHndlr))
	gohtx.AddGohtxAssetHandler()
	// The update handler
	http.Handle("/update", http.HandlerFunc(updateHndlr))

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
