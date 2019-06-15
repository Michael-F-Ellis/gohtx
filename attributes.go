package goht

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Attributes is filled in at init time with a list of attributes and the tag
// names that support them.
var Attributes map[string][]string

// checkAttr tests a single attribute name for validity in the context of a
// particular tag.
func checkAttr(tag string, a string) error {
	// data-* attributes are a special case
	if strings.HasPrefix(a, "data-") {
		name := a[5:] // everything after "data-"
		switch {
		case len(name) == 0:
			return fmt.Errorf("%s is not a valid html5 attribute", a)
		case name != strings.ToLower(name):
			return fmt.Errorf("%s: uppercase is not allowed in data-* attributes", a)
		default:
			return nil
		}
	}
	// Other attributes are validated via the Attributes map.
	tags, found := Attributes[a]
	switch {
	case !found:
		return fmt.Errorf("%s is not a valid html5 attribute", a)
	case tags[0] == "*": // global attribute
		return nil
	case !stringInSlice(tag, tags):
		return fmt.Errorf("%s is not a valid attribute for %s", a, tag)

	}
	return nil
}

// stringInSlice returns true if string s is an element in slice ss.
func stringInSlice(s string, ss []string) bool {
	for _, v := range ss {
		if s == v {
			return true
		}
	}
	return false
}

// splitOnFirstEqualSign returns the portion of a string preceding the first
// equal sign or the entire string if no equal sign is found. This function is
// not currently used in this package.
func splitOnFirstEqualSign(av string) (a string, v string) {
	sp := strings.SplitN(av, "=", 2)
	a = sp[0]
	if len(sp) == 2 {
		v = sp[1]
	}
	return
}

// checkTagAttributes builds an html string containing tag with attrs and
// parses it, calling checkAttr on each attribute found. It returns a slice of
// errors found. The slice will be empty if no errors where detected. There is
// a separate err return that should be checked. It will be nil unless the
// attrs string is so malformed that it can't be parsed.
func checkTagAttributes(tag, attrs string) (errs []error, err error) {
	// s := `<a href="foo" id="fooid" checked>Foo</a>`
	s := fmt.Sprintf("<%s %s></%s>", tag, attrs, tag)
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		return
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, a := range n.Attr {
				if err := checkAttr(tag, a.Key); err != nil {
					errs = append(errs, err)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return
}

// AttributeErrors is a struct returned by CheckAttributes. It contains the
// tag, the string of attributes and a slice of errors returned by
// checkTagAttributes.
type AttributeErrors struct {
	tag   string  // an html tag
	attrs string  // a string of zero or more attributes
	errs  []error // errors found in attr
}

// CheckAttributes walks through an ElementTree and checks each tag to verify that
// the attribute names associated with the tag are valid for that tag. It returns a slice
// AttributeErrors. The slice will be empty if no errors were found.
func (e *HtmlTree) CheckAttributes(perrs *[]AttributeErrors) {
	errslice, _ := checkTagAttributes(e.T, e.A)
	if len(errslice) != 0 {
		*perrs = append(*perrs, AttributeErrors{e.T, e.A, errslice})
	}
	for _, c := range e.C {
		switch t := c.(type) {
		case string:
			continue // nothing to do for string content
		case *HtmlTree:
			c.(*HtmlTree).CheckAttributes(perrs)
		default:
			panic(fmt.Sprintf("Don't know how to check attributes of %T", t))
		}
	}
}

func init() {
	// Derived from https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes#Attribute_list
	// Attributes marked as "legacy" in above reference are omitted here in the interest of
	// encouraging better HTML5 compliance.
	Attributes = map[string][]string{
		"accept":          []string{"form", "input"},
		"accept-charset":  []string{"form"},
		"accesskey":       []string{"*"},
		"action":          []string{"form"},
		"align":           []string{"applet", "caption", "col", "colgroup", "hr", "iframe", "img", "table", "tbody", "td", "tfoot", "th", "thead", "tr"},
		"allow":           []string{"iframe"},
		"alt":             []string{"applet", "area", "img", "input"},
		"async":           []string{"script"},
		"autocapitalize":  []string{"*"},
		"autocomplete":    []string{"form", "input", "textarea"},
		"autofocus":       []string{"button", "input", "keygen", "select", "textarea"},
		"autoplay":        []string{"audio", "video"},
		"buffered":        []string{"audio", "video"},
		"challenge":       []string{"keygen"},
		"charset":         []string{"meta", "script"},
		"checked":         []string{"command", "input"},
		"cite":            []string{"blockquote", "del", "ins", "q"},
		"class":           []string{"*"},
		"code":            []string{"applet"},
		"codebase":        []string{"applet"},
		"cols":            []string{"textarea"},
		"colspan":         []string{"td", "th"},
		"content":         []string{"meta"},
		"contenteditable": []string{"*"},
		"contextmenu":     []string{"*"},
		"controls":        []string{"audio", "video"},
		"coords":          []string{"area"},
		"crossorigin":     []string{"audio", "img", "link", "script", "video"},
		"csp":             []string{"iframe"},
		"data":            []string{"object"},
		"datetime":        []string{"del", "ins", "time"},
		"decoding":        []string{"img"},
		"default":         []string{"track"},
		"defer":           []string{"script"},
		"dir":             []string{"*"},
		"dirname":         []string{"input", "textarea"},
		"disabled":        []string{"button", "command", "fieldset", "input", "keygen", "optgroup", "option", "select", "textarea"},
		"download":        []string{"a", "area"},
		"draggable":       []string{"*"},
		"dropzone":        []string{"*"},
		"enctype":         []string{"form"},
		"for":             []string{"label", "output"},
		"form":            []string{"button", "fieldset", "input", "keygen", "label", "meter", "object", "output", "progress", "select", "textarea"},
		"formaction":      []string{"input", "button"},
		"headers":         []string{"td", "th"},
		"height":          []string{"canvas", "embed", "iframe", "img", "input", "object", "video"},
		"hidden":          []string{"*"},
		"high":            []string{"meter"},
		"href":            []string{"a", "area", "base", "link"},
		"hreflang":        []string{"a", "area", "link"},
		"http-equiv":      []string{"meta"},
		"icon":            []string{"command"},
		"id":              []string{"*"},
		"importance":      []string{"iframe", "img", "link", "script"},
		"integrity":       []string{"link", "script"},
		"ismap":           []string{"img"},
		"itemprop":        []string{"*"},
		"keytype":         []string{"keygen"},
		"kind":            []string{"track"},
		"label":           []string{"track"},
		"lang":            []string{"*"},
		"language":        []string{"script"},
		"lazyload":        []string{"img", "iframe"},
		"list":            []string{"input"},
		"loop":            []string{"audio", "bgsound", "marquee", "video"},
		"low":             []string{"meter"},
		"manifest":        []string{"html"},
		"max":             []string{"input", "meter", "progress"},
		"maxlength":       []string{"input", "textarea"},
		"minlength":       []string{"input", "textarea"},
		"media":           []string{"a", "area", "link", "source", "style"},
		"method":          []string{"form"},
		"min":             []string{"input", "meter"},
		"multiple":        []string{"input", "select"},
		"muted":           []string{"audio", "video"},
		"name":            []string{"button", "form", "fieldset", "iframe", "input", "keygen", "object", "output", "select", "textarea", "map", "meta", "param"},
		"novalidate":      []string{"form"},
		"open":            []string{"details"},
		"optimum":         []string{"meter"},
		"pattern":         []string{"input"},
		"ping":            []string{"a", "area"},
		"placeholder":     []string{"input", "textarea"},
		"poster":          []string{"video"},
		"preload":         []string{"audio", "video"},
		"radiogroup":      []string{"command"},
		"readonly":        []string{"input", "textarea"},
		"rel":             []string{"a", "area", "link"},
		"required":        []string{"input", "select", "textarea"},
		"reversed":        []string{"ol"},
		"rows":            []string{"textarea"},
		"rowspan":         []string{"td", "th"},
		"sandbox":         []string{"iframe"},
		"scope":           []string{"th"},
		"scoped":          []string{"style"},
		"selected":        []string{"option"},
		"shape":           []string{"a", "area"},
		"size":            []string{"input", "select"},
		"sizes":           []string{"link", "img", "source"},
		"slot":            []string{"*"},
		"span":            []string{"col", "colgroup"},
		"spellcheck":      []string{"*"},
		"src":             []string{"audio", "embed", "iframe", "img", "input", "script", "source", "track", "video"},
		"srcdoc":          []string{"iframe"},
		"srclang":         []string{"track"},
		"srcset":          []string{"img", "source"},
		"start":           []string{"ol"},
		"step":            []string{"input"},
		"style":           []string{"*"},
		"summary":         []string{"table"},
		"tabindex":        []string{"*"},
		"target":          []string{"a", "area", "base", "form"},
		"title":           []string{"*"},
		"translate":       []string{"*"},
		"type":            []string{"button", "input", "command", "embed", "object", "script", "source", "style", "menu"},
		"usemap":          []string{"img", "input", "object"},
		"value":           []string{"button", "option", "input", "li", "meter", "progress", "param"},
		"width":           []string{"canvas", "embed", "iframe", "img", "input", "object", "video"},
		"wrap":            []string{"textarea"},
	}
}
