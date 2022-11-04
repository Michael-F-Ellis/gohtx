package main

import (
	"testing"
)

func TestEval(t *testing.T) {
	code := `htx := Html("",Head("",Body("","hello")))`
	want := "\n<html>\n  <head>\n    <body>hello\n    </body>\n  </head>\n</html>\n"
	got := eval(code)
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
