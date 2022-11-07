// The gohtx app skeleton demonstrates the essentials of using gohtx, htmx,
// hyperscript and bulma to build a single binary containing a server and
// embedded content to generate and serve a responsive application.
package main

import (
	"embed"
	"flag"
	"log"
)

//go:embed fragments/*.txt
var fragments embed.FS // Small examples for the playground

var (
	// The following are set by flag.Parse from command line arguments.
	HostPort    string // e.g. localhost:8080 or secure.example.com:443
	CertPath    string // Path to a valid cert file.
	CertKeyPath string // Path to a valid cert key file
	Fragments   map[string]string
)

func init() {
	var err error
	Fragments, err = getFragments()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println("Starting the server")
	flag.StringVar(&HostPort, "p", "localhost:8080", `hostname (or IP) and port to serve on.`)
	flag.StringVar(&CertPath, "c", "", `path to a valid certificate file`)
	flag.StringVar(&CertKeyPath, "k", "", `path to a valid certificate key file`)
	flag.Parse()
	Serve()
}
