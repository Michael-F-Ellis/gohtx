package gohtx

import (
	"testing"

	"github.com/go-test/deep"
)

func TestGohtify(t *testing.T) {
	type testcase struct {
		html   string
		exp    string
		fmtexp string
	}
	tcases := []testcase{
		// Bare comment
		{`<!--This is a comment-->`,
			"Comment(`This is a comment`)",
			"Comment(`This is a comment`)"},

		// Comment child node
		{`<div><!--This is a comment--></div>`,
			"Div(``,Comment(`This is a comment`))",
			"Div(``,\n\tComment(`This is a comment`))"},

		// empty tag type
		{`<br>`, "Br(``)",
			"Br(``)"},

		// empty tag surrounded by whitespace
		{"\t<br>\n ", "Br(``)",
			"Br(``)"},

		// tag with attribute
		{`<br id="foo">`, "Br(`id=\"foo\"`)",
			"Br(`id=\"foo\"`)"},

		//more than one attribute
		{`<br id="foo" class="bar bare">`,
			"Br(`id=\"foo\" class=\"bar bare\"`)",
			"Br(`id=\"foo\" class=\"bar bare\"`)"},

		// non-empty tag type
		{`<div>hello</div>`, "Div(``,`hello`)",
			"Div(``, `hello`)"},

		// child elements including plain text
		{`<div><p>hello</p>xyz</div>`,
			"Div(``,P(``,`hello`),`xyz`)",
			"Div(``,\n\tP(``, `hello`),\n\t`xyz`)"},

		// child elements
		{`<div><p>hello</p><p>bye</p></div>`,
			"Div(``,P(``,`hello`),P(``,`bye`))",
			"Div(``,\n\tP(``, `hello`),\n\tP(``, `bye`))"},

		// whitespace before child tag
		{"<div>\n<p>hola</p></div>",
			"Div(``,P(``,`hola`))",
			"Div(``,\n\tP(``, `hola`))"},

		// child element with attribute
		{`<div><p class="bar">hello</p></div>`,
			"Div(``,P(`class=\"bar\"`,`hello`))",
			"Div(``,\n\tP(`class=\"bar\"`, `hello`))"},

		// Two nodes with no outer element
		{`<div>hi</div><div>bye</div>`,
			"Null(Div(``,`hi`),Div(``,`bye`))",
			"Null(\n\tDiv(``, `hi`),\n\tDiv(``, `bye`))",
		},
	}
	ignore := map[string]struct{}{"html": {}, "head": {}, "body": {}}
	for _, tc := range tcases {
		var got string
		err := Gohtify(tc.html, false, ignore, &got)
		if err != nil {
			t.Errorf("%v", err)
			continue
		}
		if diff := deep.Equal(got, tc.exp); diff != nil {
			t.Errorf("%v", diff)
			continue
		}
	}
	for _, tc := range tcases {
		var got string
		err := Gohtify(tc.html, true, ignore, &got)
		if err != nil {
			t.Errorf("%v", err)
			continue
		}
		if diff := deep.Equal(got, tc.fmtexp); diff != nil {
			t.Errorf("\n%v", diff)
			continue
		}
	}
}
