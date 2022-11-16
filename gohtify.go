package gohtx

import (
	"fmt"
	"go/format"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Gohtify parses an html string, htext, and returns equivalent goht code in
// gohttext. The ignoreTags map contains one or more tags to be ignored as map
// keys. These are typically "html", "head" and "body" because html.Parse
// creates them if they're not present in the string. That's usually not desired
// since the chief use of Gohtify is to turn useful fragments of html into
// into equivalent Go code.
func Gohtify(htext string, gofmt bool, ignoreTags map[string]struct{}, gohttext *string) (err error) {
	// parse with net/html package
	htext = strings.TrimSpace(htext)
	doc, err := html.ParseFragment(strings.NewReader(htext), nil)
	if err != nil {
		return
	}
	// Define a func that will be called to walk the parsed node tree contained
	// in doc
	var (
		f         func(*html.Node)
		depth     int // depth of recursion
		zeroCount int // number of times depth returns to zero
	)
	f = func(n *html.Node) {
		var ignore bool
		switch n.Type {
		case html.ElementNode:
			_, ignore = ignoreTags[n.Data]
			if !ignore {
				// Get the tag name, capitalize the first letter and
				// append an opening parenthesis, and add the attributes as a
				// back quoted string, e.g. "<div id=foo" -> "Div(`id=foo`"
				*gohttext += cases.Title(language.Und).String(n.Data) + "(`" + nodeAttrs(n.Attr) + "`"
				maybeAddComma(n, gohttext, gofmt)
			}
			// recurse on the content, if any
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if !ignore {
					depth++
				}
				f(c)
				if !ignore {
					depth--
				}
			}
			// close the function call
			if !ignore {
				*gohttext += ")"
				if depth == 0 {
					zeroCount++
				}
				// If there's more than one node and this isn't the last, append either
				// a comma or a comma+newline depending the value of gofmt.
				if n.NextSibling != nil {
					switch gofmt {
					case true:
						*gohttext += ",\n"
					case false:
						*gohttext += ","
					}
				}

			}
		case html.TextNode:
			// if the content is a string, strip leading and trailing
			// whitespace, then append it enclosed in back quotes unless it
			// consists entirely of whitespace. In that case, omit the back
			// quotes.
			switch isWhiteSpace(n.Data) {
			case true:
				*gohttext += ""
			case false:
				*gohttext += "`" + strings.TrimSpace(n.Data) + "`"
			}
		case html.CommentNode:
			*gohttext += "Comment(`" + strings.TrimSpace(n.Data) + "`)"
		}
	}
	// Walk the tree
	for _, n := range doc {
		f(n)
	}
	// remove trailing comma or comma+newline
	if gofmt {
		// *gohttext = strings.TrimSuffix(*gohttext, ",\n")
		if zeroCount > 1 {
			*gohttext = "Null(\n" + *gohttext + ")"
		}
		var buf []byte
		buf, err = format.Source([]byte(*gohttext))
		if err != nil {
			err = fmt.Errorf("got error: %v trying to fmt %s", err, *gohttext)
			return
		}
		*gohttext = string(buf)
	} else {
		// *gohttext = strings.TrimSuffix(*gohttext, ",")
		if zeroCount > 1 {
			*gohttext = "Null(" + *gohttext + ")"
		}
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

// maybeAddComma appends a comma if the first child node is not nil and is not a
// TextNode containing only whitespace and having no siblings. If fmt is true
// and a comma was added, it will also add a newline unless the first child is a
// non-whitespace TextNode.
func maybeAddComma(n *html.Node, gohtText *string, fmt bool) {
	switch {
	case n.FirstChild == nil:
		// e.g. '<div></div>' --> Div(``,)
		return // don't add a comma or newline
	case n.FirstChild.Type == html.TextNode:
		if isWhiteSpace(n.FirstChild.Data) {
			switch {
			case n.FirstChild == n.LastChild:
				// e.g. '<div>   </div>' --> Div(``,
				return // don't add a comma or newline
			default:
				// e.g . '<div>   <p>hello</p></div>' --> Div(``,\n
				*gohtText += ","
				if fmt {
					*gohtText += "\n"
				}
				return
			}
		}
		// if not whitespace, add a comma but not a newline
		// e.g. <div>hello</div> --> Div(``,
		*gohtText += ","
		return

	default:
		*gohtText += ","
		if fmt {
			*gohtText += "\n"
		}
	}
}
