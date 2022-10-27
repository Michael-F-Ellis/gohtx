// The gohtx app skeleton
package main

import (
	"flag"
	"log"
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
