// gohtify reads html fragments from stdin and emits goht code on stdout.
package main

import (
	"fmt"
	"io"
	"os"

	"github.com/Michael-F-Ellis/gohtx"
)

func main() {
	ignoreTags := map[string]struct{}{
		"html": {}, "head": {}, "body": {},
	}
	html, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var gohtText string
	err = gohtx.Gohtify(string(html), ignoreTags, &gohtText)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(gohtText)
	os.Exit(0)
}
