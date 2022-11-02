package gohtx

import (
	"bytes"
	"strings"
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
func TestCustomHeadContent(t *testing.T) {
	var buf bytes.Buffer
	// test all 8 possible combinations
	h000 := Head(``, CustomHeadContent(false, false, false))
	h001 := Head(``, CustomHeadContent(false, false, true))
	h010 := Head(``, CustomHeadContent(false, true, false))
	h011 := Head(``, CustomHeadContent(false, true, true))
	h100 := Head(``, CustomHeadContent(true, false, false))
	h101 := Head(``, CustomHeadContent(true, false, true))
	h110 := Head(``, CustomHeadContent(true, true, false))
	h111 := Head(``, CustomHeadContent(true, true, true))
	err := Render(h000, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s := buf.String()
	if strings.Contains(s, "htmx.min.js") {
		t.Errorf("unexpected %s in \n%s\n", "htmx", s)
	}
	if strings.Contains(s, "hyperscript.js") {
		t.Errorf("unexpected %s in \n%s\n", "hyperscript.js", s)
	}
	if strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("unexpected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

	buf.Reset()
	err = Render(h001, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s = buf.String()
	if strings.Contains(s, "htmx.min.js") {
		t.Errorf("unexpected %s in \n%s\n", "htmx", s)
	}
	if strings.Contains(s, "hyperscript.js") {
		t.Errorf("unexpected %s in \n%s\n", "hyperscript.js", s)
	}
	if !strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("expected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

	buf.Reset()
	err = Render(h010, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s = buf.String()
	if strings.Contains(s, "htmx.min.js") {
		t.Errorf("unexpected %s in \n%s\n", "htmx", s)
	}
	if !strings.Contains(s, "hyperscript.js") {
		t.Errorf("expected %s in \n%s\n", "hyperscript.js", s)
	}
	if strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("unexpected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

	buf.Reset()
	err = Render(h011, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s = buf.String()
	if strings.Contains(s, "htmx.min.js") {
		t.Errorf("unexpected %s in \n%s\n", "htmx", s)
	}
	if !strings.Contains(s, "hyperscript.js") {
		t.Errorf("expected %s in \n%s\n", "hyperscript.js", s)
	}
	if !strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("expected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

	buf.Reset()
	err = Render(h100, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s = buf.String()
	if !strings.Contains(s, "htmx.min.js") {
		t.Errorf("expected %s in \n%s\n", "htmx", s)
	}
	if strings.Contains(s, "hyperscript.js") {
		t.Errorf("unexpected %s in \n%s\n", "hyperscript.js", s)
	}
	if strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("unexpected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

	buf.Reset()
	err = Render(h101, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s = buf.String()
	if !strings.Contains(s, "htmx.min.js") {
		t.Errorf("expected %s in \n%s\n", "htmx", s)
	}
	if strings.Contains(s, "hyperscript.js") {
		t.Errorf("unexpected %s in \n%s\n", "hyperscript.js", s)
	}
	if !strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("expected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

	buf.Reset()
	err = Render(h110, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s = buf.String()
	if !strings.Contains(s, "htmx.min.js") {
		t.Errorf("expected %s in \n%s\n", "htmx", s)
	}
	if !strings.Contains(s, "hyperscript.js") {
		t.Errorf("expected %s in \n%s\n", "hyperscript.js", s)
	}
	if strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("unexpected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

	buf.Reset()
	err = Render(h111, &buf, 0)
	if err != nil {
		t.Error(err)
	}
	s = buf.String()
	if !strings.Contains(s, "htmx.min.js") {
		t.Errorf("expected %s in \n%s\n", "htmx", s)
	}
	if !strings.Contains(s, "hyperscript.js") {
		t.Errorf("expected %s in \n%s\n", "hyperscript.js", s)
	}
	if !strings.Contains(s, "bulma/css/bulma.min.css") {
		t.Errorf("expected %s in \n%s\n", "bulma/css/bulma.min.css", s)
	}

}
