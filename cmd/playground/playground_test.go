package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Michael-F-Ellis/gohtx"
)

func TestEval(t *testing.T) {
	code := `htx := Html("",Head("",Body("","hello")))`
	want := "\n<html>\n  <head>\n    <body>hello\n    </body>\n  </head>\n</html>\n"
	got, _ := eval(code, true)
	if strings.Contains(got, want) {
		t.Errorf("got %q, want %q", got, want)
		return
	}
	// TODO: Verify that result compiles without errors.
}

func TestFragments(t *testing.T) {
	placeholder, err := fragments.ReadFile("fragments/notification.txt")
	if err != nil {
		t.Errorf("%v", err)
	}
	textAreaAttrs := `class="textarea" name="code"`
	code := gohtx.Textarea(textAreaAttrs, string(placeholder))
	var buf bytes.Buffer
	err = gohtx.Render(code, &buf, 0)
	if err != nil {
		t.Errorf("%v", err)
	}
}
func TestGetFragments(t *testing.T) {
	m, err := getFragments()
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	_, ok := m["Notification"] // Should be present
	if !ok {
		t.Errorf("Notification not found")
	}
}

func TestMkSelect(t *testing.T) {
	names := []string{"one", "two"}
	s := mkSelect(names, "/url", "which", "#id")
	exp := `<div class="select is-link">` +
		`<select name="which" hx-get="/url" hx-target="#id">` +
		`<option value="one">one</option>` +
		`<option value="two">two</option>` +
		`</select></div>`
	var buf bytes.Buffer
	err := gohtx.Render(s, &buf, -1)
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	if buf.String() != exp {
		t.Errorf("\ngot %s \nexpected %s", buf.String(), exp)
	}
}
