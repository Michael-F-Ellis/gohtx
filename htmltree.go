// Atempting to re-create my Python htmltree in go.
package htmltree

import (
	"bytes"
	"fmt"
)

type HtmlTree struct {
	T     string
	A     string
	C     []interface{}
	empty bool // set to true for empty tags like <br>
}

func Render(h *HtmlTree, b *bytes.Buffer, nindent int) (err error) {
	// render the opening tag
	indent := indentation(nindent)
	b.WriteString(indent)
	b.WriteString("<")
	b.WriteString(h.T)
	// render the attributes
	if len(h.A) > 0 {
		b.WriteString(" ")
	}
	b.WriteString(h.A)
	// close the opening tag
	b.WriteString(">")

	// indentation for nested content.
	rindent := nindent
	if nindent >= 0 {
		rindent = nindent + 1
	}
	if h.empty {
		if len(h.C) == 0 {
			return nil
		} else {
			return fmt.Errorf("%s : empty tag may not have content", h.T)
		}
	}
	// otherwise, recursively render the content
	for _, c := range h.C {
		switch c.(type) {
		case string:
			b.WriteString(c.(string))
		case *HtmlTree:
			err = Render(c.(*HtmlTree), b, rindent)
			if err != nil {
				return fmt.Errorf("%s : %v", h.T, err)
			}
		default:
			return fmt.Errorf("Bad content %v. Can't render type %T! ", h.C, c)
		}
	}
	// render the closing tag
	b.WriteString(indent)
	b.WriteString("</")
	b.WriteString(h.T)
	b.WriteString(">")

	return
}

// indentation returns a string like "\n  " where the number of spaces is n * 2
// if n is 0 or greater. If n is negative, indentation returns an empty string.
// The negative case supports rendering an entire tree without newlines or
// leading spaces.
func indentation(n int) string {
	if n < 0 {
		return "" // no indentation
	}
	s := "\n"
	for i := 0; i < 2*n; i++ {
		s += " "
	}
	return s
}
