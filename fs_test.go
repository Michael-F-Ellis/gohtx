package gohtx

import (
	"bytes"
	"testing"
)

func TestGohtxFS(t *testing.T) {
	fs := GohtxFS()
	for _, fname := range []string{"htmx.min.js", "bulma/css/bulma.css"} {
		buf, err := fs.ReadFile(fname)
		if err != nil {
			t.Errorf("failed to read %s: %v", fname, err)
			continue
		}

		if len(buf) == 0 {
			t.Errorf("read %s, but it is empty", fname)
		}
	}
}

func TestDefaultHeadContent(t *testing.T) {
	var buf bytes.Buffer
	head := Head(``, DefaultHeadContent(), Title(``, "The Title"))
	err := Render(head, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	// rendered := buf.String()
	// t.Error(rendered)
}
