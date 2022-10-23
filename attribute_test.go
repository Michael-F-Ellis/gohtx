package gohtx

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
)

func TestStringInSlice(t *testing.T) {
	type items struct {
		s   string   // what to find
		ss  []string // where to look
		exp bool     // expected result
	}
	table := []items{
		{"foo", []string{"bar", "foo"}, true},
		{"foo", []string{"bar", "fool"}, false},
	}
	for _, test := range table {
		found := stringInSlice(test.s, test.ss)
		if found != test.exp {
			t.Errorf("Searched for %s. Expected %v got %v.", test.s, test.exp, found)
		}
	}
}

func TestCheckAttr(t *testing.T) {
	type items struct {
		a   string // attr
		tag string //tag
		exp error  // expected result
	}
	table := []items{
		{"href", "a", nil},
		{"hx-post", "button", nil},
		{"hx-morble", "button", errors.New("hx-morble is not a valid html5 attribute")},
		{"HREF", "a", errors.New("HREF is not a valid html5 attribute")},
		{"href", "body", errors.New("href is not a valid attribute for body")},
		{"junk", "body", errors.New("junk is not a valid html5 attribute")},
		{"data-foobar", "body", nil},
		{"data-x", "body", nil},
		{"data-hx-post", "button", nil},
		{"data-hx-morble", "button", errors.New("data-hx-morble doesn't match any valid htmx attribute")},
		{"data-fooBar", "body", errors.New("data-fooBar: uppercase is not allowed in data-* attributes")},
		{"data-", "body", errors.New("data- is not a valid html5 attribute")},
	}
	for _, test := range table {
		err := checkAttr(test.tag, test.a)
		switch err {
		case nil:
			if test.exp != nil {
				t.Errorf("Expected %v got nil.", test.exp)
			}
		default:
			if err.Error() != test.exp.Error() {
				t.Errorf("Expected %v got %v.", test.exp, err)
			}
		}
	}
}

func BenchmarkCheckAttr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkAttr("a", "href")
	}
}

func BenchmarkCheckAttrErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = checkAttr("body", "href") // invalid attribute for body
	}
}
func TestCheckTagAttributes(t *testing.T) {
	type items struct {
		tag string  //tag
		a   string  // attrs
		exp []error // expected result
	}
	table := []items{
		{"a", `href="https://example.com/foo" id="foolink"`, nil},
		{"body", `href="https://example.com/foo" id="foolink"`, []error{errors.New("href is not a valid attribute for body")}},
	}
	for _, test := range table {
		errs, err := checkTagAttributes(test.tag, test.a)
		switch err {
		case nil:
			if diff := deep.Equal(test.exp, errs); diff != nil {
				t.Errorf("%v", diff)
			}
		default:
			t.Errorf("html.Parse error: %v", err)
		}
	}
}

func TestCheckAttributes(t *testing.T) {
	type items struct {
		e   *HtmlTree
		exp []AttributeErrors
	}
	table := []items{
		{Html(""), []AttributeErrors{}},
		{Html("", "Just a string."), []AttributeErrors{}},
		{Html("", Div(`id="1"`, "Just a string.", Br(""))), []AttributeErrors{}},
	}
	for _, test := range table {
		perrs := &[]AttributeErrors{}
		test.e.CheckAttributes(perrs)
		if diff := deep.Equal(test.exp, *perrs); diff != nil {
			t.Errorf("%v", diff)
		}
	}
}

func BenchmarkCheckAttributes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = checkTagAttributes("a", `href="https://example.com/foo" id="foolink"`)
	}
}
func BenchmarkCheckTagAttributesErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = checkTagAttributes("body", `href="https://example.com/foo" id="foolink"`)
	}
}
func BenchmarkCheckAttributesErr(b *testing.B) {
	meta := Meta(`title="Demo"`)
	head := Head("id=2 class=foo", meta)
	body := Body("id=3 class=bar", Div("", "hello", Br(``)))
	html := Html("", head, body)
	perrs := &[]AttributeErrors{}
	for i := 0; i < b.N; i++ {
		html.CheckAttributes(perrs)
	}
}
