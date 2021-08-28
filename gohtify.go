package goht

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

// Gohtify parses an html string, htext, and returns equivalent goht code in
// gohttext. The ignoreTags map contains one or more tags to be ignored as map
// keys. These are typically "html", "head" and "body" because html.Parse
// creates them if they're not present in the string. That's usually not desired
// since the chief use of Gohtify is to turn useful fragments of html into
// into equivalent Go code.
func Gohtify(htext string, ignoreTags map[string]struct{}, gohttext *string) (err error) {
	// parse with net/html package
	htext = strings.TrimSpace(htext)
	doc, err := html.ParseFragment(strings.NewReader(htext), nil)
	if err != nil {
		return
	}

	// Define a func that will be called to walk the parsed node tree contained
	// in doc
	var f func(*html.Node)
	f = func(n *html.Node) {
		var ignore bool
		switch n.Type {
		case html.ElementNode:
			_, ignore = ignoreTags[n.Data]
			if !ignore {
				// Get the tag name, capitalize the first letter and
				// append an opening parenthesis, and add the attributes as a
				// back quoted string, e.g. "<div id=foo" -> "Div(`id=foo`"
				*gohttext += strings.Title(n.Data) + "(`" + nodeAttrs(n.Attr) + "`"
			}
			// recurse on the content, if any
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if !ignore {
					maybeAddComma(c, gohttext)
				}
				f(c)
			}
			// close the function call
			if !ignore {
				*gohttext += ")"
			}
		case html.TextNode:
			// if the content is a string, append it enclosed in back quotes
			// unless it consists enttirely of whitespace. In that case, omit
			// the back quotes.
			switch isWhiteSpace(n.Data) {
			case true:
				*gohttext += n.Data
			case false:
				*gohttext += "`" + n.Data + "`"
			}
		}

	}
	// Walk the tree
	for _, n := range doc {
		f(n)
	}
	return
}
func nodeAttrs(attributes []html.Attribute) string {
	attrs := []string{}
	for _, a := range attributes {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	return strings.Join(attrs, " ")
}

// isWhiteSpace returns true if the string contains only whitespace characters,
// false otherwise.
func isWhiteSpace(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// maybeAddComma appends a comma if the previous child node is not nil and is
// not a TextNode containing only whitespace.
func maybeAddComma(n *html.Node, gohtText *string) {
	if n.PrevSibling == nil {
		*gohtText += ","
		return
	}
	switch n.PrevSibling.Type {
	case html.TextNode:
		if !isWhiteSpace(n.PrevSibling.Data) {
			*gohtText += ","
		}
	default:
		*gohtText += ","
	}
}
