// Atempting to re-create my Python htmltree in go.
package gohtx

import (
	"bytes"
	"fmt"
	"strings"
)

// HtmlTree represents a tree of html content.
type HtmlTree struct {
	T     string        // html tagname, e.g. 'head'
	A     string        // zero or more html attributes, e.g 'id=1 class="foo"'
	C     []interface{} // a slice of content whose elements may be strings or *HtmlTree
	empty bool          // set to true for empty tags like <br>
}

// Render walks through HtmlTree h and writes html text to byte buffer b. The
// nindent argument specifies whether and how much to indent the output where -1
// means render all on one line and 0 means indent each inner tag by 2 spaces
// relative to its parent. Render returns an error when it encounters invalid
// content.
func Render(h *HtmlTree, b *bytes.Buffer, nindent int) (err error) {
	// render the opening tag unless it's the Null tag
	indent := indentation(nindent)
	if h.T != "null" {
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
	}

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
		switch c := c.(type) {
		case string:
			b.WriteString(c)
		case *HtmlTree:
			err = Render(c, b, rindent)
			if err != nil {
				return fmt.Errorf("%s : %v", h.T, err)
			}
		default:
			return fmt.Errorf("Bad content %v. Can't render type %T! ", h.C, c)
		}
	}
	// render the closing tag unless it's the Null tag
	if h.T == "null" {
		return // nothing more to write
	}
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

// Ids returns a slice of string containing all the id attributes found in tree.
// It will return an error if the search finds a malformed id, multiple ids in
// the same tag or the same id in different tags.
func Ids(tree *HtmlTree, ids *[]string) (err error) {
	// split the attributes string on whitespace
	attrs := strings.Fields(tree.A)
	// search for strings starting with "id="
	var n int // number of id strings found
	var id string
	for _, a := range attrs {
		if strings.HasPrefix(strings.ToLower(a), "id=") {
			id = a[3:]
			if len(id) == 0 {
				err = fmt.Errorf(`empty id attribute in '%s'`, tree.A)
				return
			}
			n++
		}
	}
	// if there are more than one, return an error
	if n > 1 {
		err = fmt.Errorf("more than one id attribute in '%s' (attributes of %s).", tree.A, tree.T)
		return
	}
	if n == 1 {
		*ids = append(*ids, strings.Trim(id, `"'`))
	}
	// recurse over the tree content
	for _, c := range tree.C {
		switch c := c.(type) {
		case string: // terminal node for this branch
			continue
		case *HtmlTree:
			err = Ids(c, ids)
		default:
			err = fmt.Errorf("Bad content %v. Can't render type %T! ", tree.C, c)
			return
		}
	}
	// Test for duplicates using map insertion
	var m = make(map[string]int)
	for i, s := range *ids {
		_, exist := m[s]
		if exist {
			err = fmt.Errorf("duplicated id %s", s)
			break
		}
		m[s] = i
	}
	return
}
