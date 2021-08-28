package goht

import (
	"testing"

	"github.com/go-test/deep"
)

func TestGohtify(t *testing.T) {
	type testcase struct {
		html string
		exp  string
	}
	tcases := []testcase{
		// empty tag type
		{`<br>`, "Br(``)"},
		// empty tag surrounded by whitespace
		{"\t<br>\n ", "Br(``)"},
		// tag with attribute
		{`<br id="foo">`, "Br(`id=\"foo\"`)"},
		//more than one attribute
		{`<br id="foo" class="bar bare">`,
			"Br(`id=\"foo\" class=\"bar bare\"`)"},
		// non-empty tag type
		{`<div>hello</div>`, "Div(``,`hello`)"},
		// child elements including plain text
		{`<div><p>hello</p>xyz</div>`,
			"Div(``,P(``,`hello`),`xyz`)"},
		// child elements
		{`<div><p>hello</p><p>bye</p></div>`,
			"Div(``,P(``,`hello`),P(``,`bye`))"},
		// whitespace before child tag
		{"<div>\n<p>hello</p></div>",
			"Div(``,\nP(``,`hello`))"},
		// child element with attribute
		{`<div><p class="bar">hello</p></div>`,
			"Div(``,P(`class=\"bar\"`,`hello`))"},
	}
	ignore := map[string]struct{}{"html": {}, "head": {}, "body": {}}
	for _, tc := range tcases {
		var got string
		err := Gohtify(tc.html, ignore, &got)
		if err != nil {
			t.Errorf("%v", err)
			continue
		}
		if diff := deep.Equal(got, tc.exp); diff != nil {
			t.Errorf("%v", diff)
			continue
		}
	}
}
