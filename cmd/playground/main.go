// The gohtx app skeleton demonstrates the essentials of using gohtx, htmx,
// hyperscript and bulma to build a single binary containing a server and
// embedded content to generate and serve a responsive application.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/bitfield/script"
)

var (
	// The following are set by flag.Parse from command line arguments.
	HostPort    string // e.g. localhost:8080 or secure.example.com:443
	CertPath    string // Path to a valid cert file.
	CertKeyPath string // Path to a valid cert key file.
)

func main() {
	log.Println("Starting the server")
	flag.StringVar(&HostPort, "p", "localhost:8080", `hostname (or IP) and port to serve on.`)
	flag.StringVar(&CertPath, "c", "", `path to a valid certificate file`)
	flag.StringVar(&CertKeyPath, "k", "", `path to a valid certificate key file`)
	flag.Parse()
	Serve()
}

// eval is called to evaluate Go code entered into the playground.
func eval(input string) (htm string) {
	code := fmt.Sprintf(template, input)
	tmpfile, err := ioutil.TempFile("temp", "*.go")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpfile.Name())
	if _, err := tmpfile.Write([]byte(code)); err != nil {
		log.Fatal(err)
	}
	htm, err = script.Exec("go run " + tmpfile.Name()).String()
	if err != nil {
		log.Fatal(err)
	}
	return
}

var template = `
package main

import (
	"bytes"
	"fmt"

	. "github.com/Michael-F-Ellis/gohtx"
)

func main() {

	%s

	var buf bytes.Buffer
	err := Render(htx, &buf, 0)
	if err != nil {
		buf.Reset()
		err = Render(P("", "render failed: "+err.Error()), &buf, 0)
		// This should never happen ...
		if err != nil {
			panic(err)
		}
	}
	fmt.Println(buf.String())
}`
