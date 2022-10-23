package gohtx

import "testing"

func TestGohtxFS(t *testing.T) {
	fs := GohtxFS()
	for _, fname := range []string{"htmx.min.js", "bulma.css"} {
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
