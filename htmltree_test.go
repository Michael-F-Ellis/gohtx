package gohtx

import (
	"bytes"
	"testing"

	"github.com/go-test/deep"
)

func TestRender(t *testing.T) {
	type items struct {
		e   *HtmlTree //what to render
		exp string    // expected result
	}
	table := []items{
		{Comment(`This is a comment`), "<!-- This is a comment-->"},
		{Html(""), "<html></html>"},
		{P(`class=myclass`), "<p class=myclass></p>"},
		{P(`data-foo="foo text"`), `<p data-foo="foo text"></p>`},
		{Br(``), `<br>`},
		{Null(``), ``},
		{Null(Br(``), Br(``)), `<br><br>`},
	}
	for _, test := range table {
		var b bytes.Buffer
		err := Render(test.e, &b, -1)
		if err != nil {
			t.Errorf("Render failed: %v", err)
		}
		r := b.String()
		if r != test.exp {
			t.Errorf("Expected %s, got %s", test.exp, r)
		}
	}
}

func BenchmarkRender(b *testing.B) {
	meta := Meta(`title="Demo"`)
	head := Head("id=2 class=foo", meta)
	body := Body("id=3 class=bar", Div("", "hello", Br(``)))
	html := Html("", head, body)

	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		_ = Render(html, &b, -1)
	}
}
func TestIds(t *testing.T) {
	type testcase struct {
		tree *HtmlTree
		exp  []string
	}
	// These cases should not return errors
	tcases := []testcase{
		// No id
		{A(`href=something`, "foo"), []string{}},
		// One element, no content
		{Dialog(`id=myid`, "This is a useless dialog."), []string{"myid"}},
		// Quoted id
		{Dialog(`id="myid"`, "This is a useless dialog."), []string{"myid"}},
		// nested content
		{Div(`id=divid`, P(`class=foo id=pid`, "")), []string{"divid", "pid"}},
	}

	for _, tc := range tcases {
		ids := []string{}
		err := Ids(tc.tree, &ids)
		if err != nil {
			t.Errorf("%v", err)
			continue
		}
		if diff := deep.Equal(ids, tc.exp); diff != nil {
			t.Errorf("%v", diff)
			continue
		}
	}
	type etestcase struct {
		tree *HtmlTree
	}

	// These cases should return errors
	ecases := []etestcase{
		// more than one id
		{Div(`id=foo class=x id=bar`, "")},
		// bad content
		{Div(`id=foo`, 42)},
		// empty id
		{Div(`id= foo`, 42)},
		// duplicate ids
		{Div(`id=foo`, P(`id=foo`))},
	}
	for _, tc := range ecases {
		ids := []string{}
		err := Ids(tc.tree, &ids)
		if err == nil {
			t.Errorf("%s", "expected an error, got nil")
			continue
		}
	}
}
