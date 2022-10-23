package gohtx

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
		case strings.HasPrefix(name, "hx") && !isValidHxAttribute(name):
			return fmt.Errorf("%s doesn't match any valid htmx attribute", a)
		default:
			return nil
		}
	}
	if isValidHxAttribute(a) {
		return nil
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
func isValidHxAttribute(a string) bool {
	return stringInSlice(a, HxAttrs)
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
//func splitOnFirstEqualSign(av string) (a string, v string) {
//	sp := strings.SplitN(av, "=", 2)
//	a = sp[0]
//	if len(sp) == 2 {
//		v = sp[1]
//	}
//	return
//}

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
	Tag   string  // an html tag
	Attrs string  // a string of zero or more attributes
	Errs  []error // errors found in attrs
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
		"accept":          {"form", "input"},
		"accept-charset":  {"form"},
		"accesskey":       {"*"},
		"action":          {"form"},
		"align":           {"applet", "caption", "col", "colgroup", "hr", "iframe", "img", "table", "tbody", "td", "tfoot", "th", "thead", "tr"},
		"allow":           {"iframe"},
		"alt":             {"applet", "area", "img", "input"},
		"async":           {"script"},
		"autocapitalize":  {"*"},
		"autocomplete":    {"form", "input", "textarea"},
		"autofocus":       {"button", "input", "keygen", "select", "textarea"},
		"autoplay":        {"audio", "video"},
		"buffered":        {"audio", "video"},
		"challenge":       {"keygen"},
		"charset":         {"meta", "script"},
		"checked":         {"command", "input"},
		"cite":            {"blockquote", "del", "ins", "q"},
		"class":           {"*"},
		"code":            {"applet"},
		"codebase":        {"applet"},
		"cols":            {"textarea"},
		"colspan":         {"td", "th"},
		"content":         {"meta"},
		"contenteditable": {"*"},
		"contextmenu":     {"*"},
		"controls":        {"audio", "video"},
		"coords":          {"area"},
		"crossorigin":     {"audio", "img", "link", "script", "video"},
		"csp":             {"iframe"},
		"data":            {"object"},
		"datetime":        {"del", "ins", "time"},
		"decoding":        {"img"},
		"default":         {"track"},
		"defer":           {"script"},
		"dialog":          {"*"},
		"dir":             {"*"},
		"dirname":         {"input", "textarea"},
		"disabled":        {"button", "command", "fieldset", "input", "keygen", "optgroup", "option", "select", "textarea"},
		"download":        {"a", "area"},
		"draggable":       {"*"},
		"dropzone":        {"*"},
		"enctype":         {"form"},
		"for":             {"label", "output"},
		"form":            {"button", "fieldset", "input", "keygen", "label", "meter", "object", "output", "progress", "select", "textarea"},
		"formaction":      {"input", "button"},
		"headers":         {"td", "th"},
		"height":          {"canvas", "embed", "iframe", "img", "input", "object", "video"},
		"hidden":          {"*"},
		"high":            {"meter"},
		"href":            {"a", "area", "base", "link"},
		"hreflang":        {"a", "area", "link"},
		"http-equiv":      {"meta"},
		"icon":            {"command"},
		"id":              {"*"},
		"importance":      {"iframe", "img", "link", "script"},
		"integrity":       {"link", "script"},
		"ismap":           {"img"},
		"itemprop":        {"*"},
		"keytype":         {"keygen"},
		"kind":            {"track"},
		"label":           {"track"},
		"lang":            {"*"},
		"language":        {"script"},
		"lazyload":        {"img", "iframe"},
		"list":            {"input"},
		"loop":            {"audio", "bgsound", "marquee", "video"},
		"low":             {"meter"},
		"manifest":        {"html"},
		"max":             {"input", "meter", "progress"},
		"maxlength":       {"input", "textarea"},
		"minlength":       {"input", "textarea"},
		"media":           {"a", "area", "link", "source", "style"},
		"method":          {"form"},
		"min":             {"input", "meter"},
		"multiple":        {"input", "select"},
		"muted":           {"audio", "video"},
		"name":            {"button", "form", "fieldset", "iframe", "input", "keygen", "object", "output", "select", "textarea", "map", "meta", "param"},
		"novalidate":      {"form"},
		"open":            {"details"},
		"optimum":         {"meter"},
		"pattern":         {"input"},
		"ping":            {"a", "area"},
		"placeholder":     {"input", "textarea"},
		"poster":          {"video"},
		"preload":         {"audio", "video"},
		"radiogroup":      {"command"},
		"readonly":        {"input", "textarea"},
		"rel":             {"a", "area", "link"},
		"required":        {"input", "select", "textarea"},
		"reversed":        {"ol"},
		"rows":            {"textarea"},
		"rowspan":         {"td", "th"},
		"sandbox":         {"iframe"},
		"scope":           {"th"},
		"scoped":          {"style"},
		"selected":        {"option"},
		"shape":           {"a", "area"},
		"size":            {"input", "select"},
		"sizes":           {"link", "img", "source"},
		"slot":            {"*"},
		"span":            {"col", "colgroup"},
		"spellcheck":      {"*"},
		"src":             {"audio", "embed", "iframe", "img", "input", "script", "source", "track", "video"},
		"srcdoc":          {"iframe"},
		"srclang":         {"track"},
		"srcset":          {"img", "source"},
		"start":           {"ol"},
		"step":            {"input"},
		"style":           {"*"},
		"summary":         {"table"},
		"tabindex":        {"*"},
		"target":          {"a", "area", "base", "form"},
		"title":           {"*"},
		"translate":       {"*"},
		"type":            {"button", "input", "command", "embed", "object", "script", "source", "style", "menu"},
		"usemap":          {"img", "input", "object"},
		"value":           {"button", "option", "input", "li", "meter", "progress", "param"},
		"width":           {"canvas", "embed", "iframe", "img", "input", "object", "video"},
		"wrap":            {"textarea"},
	}
}
