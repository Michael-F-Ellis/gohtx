// The gohtx app skeleton demonstrates the essentials of using gohtx, htmx,
// hyperscript and bulma to build a single binary containing a server and
// embedded content to generate and serve a responsive application.
package main

import (
	"embed"
	"flag"
	"log"
	"os"
	"strings"
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
		// Trim the first line and use the trimmed string as the key for the
		// map entry.
		code := strings.TrimSpace(string(buf))
		lines := strings.Split(code, "\n")
		name := strings.TrimSpace(strings.ReplaceAll(lines[0], "/", ""))
		m[name] = code
	}
	return
}
