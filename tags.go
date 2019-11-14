package goht

// Wrappers for html element tags.
//
// Functions are grouped in the categories given at
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element and
// are alphabetical within groups.
//
// Conventions:
//     Functions are named by tag with initial caps, e.g. Html()
//
//     The signature for non-empty tags is Tagname(a string, c ...interface{}) *ElementTree
//     The signature for empty tags is Tagname(a string) *ElementTree
//
//     Empty refers to elements that enclose no content and need no closing tag.
//
// Obsolete and Deprecated Elements:
// No pull requests will be accepted for
// acronym, applet, basefont, big, blink, center, command, content,
// dir, element, font, frame, frameset, isindex, keygen, listing,
// marquee, multicol, nextid, noembed, plaintext, shadow, spacer,
// strike, tt, xmp .
// But you can define these or any other tag quite simply by following the
// pattern used for the tags defined in this file.

// Main Root
func Html(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"html", a, c, false}
}

// Document Metadata
// TODO: base

func Head(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"head", a, c, false}
}

func Body(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"body", a, c, false}
}

func Link(a string) *HtmlTree {
	return &HtmlTree{"link", a, []interface{}{}, true}
}

func Meta(a string) *HtmlTree {
	return &HtmlTree{"meta", a, []interface{}{}, true}
}

func Title(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"title", a, c, false}
}

// Style is a special case in the sense that the only
// valid content is one or more strings of CSS. At this time
// there's no check to complain about other content.
func Style(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"style", a, c, false}
}

// interface{} Sectioning
// TODO hgroup

func Address(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"address", a, c, false}
}

func Article(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"article", a, c, false}
}

func Aside(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"aside", a, c, false}
}

func Footer(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"footer", a, c, false}
}

func Header(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"header", a, c, false}
}

func H1(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h1", a, c, false}
}

func H2(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h2", a, c, false}
}

func H3(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h3", a, c, false}
}

func H4(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h4", a, c, false}
}

func H5(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h5", a, c, false}
}

func H6(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h6", a, c, false}
}

func Nav(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"nav", a, c, false}
}

func Section(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"section", a, c, false}

}

// Text interface{}

func Blockquote(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"blockquote", a, c, false}
}

func Dd(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"dd", a, c, false}
}

func Div(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"div", a, c, false}
}

func Dl(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"dl", a, c, false}
}

func Dt(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"dt", a, c, false}
}

func Figcaption(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"figcaption", a, c, false}
}

func Figure(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"figure", a, c, false}
}

func Hr(a string) *HtmlTree {
	return &HtmlTree{"hr", a, []interface{}{}, true}
}

func Li(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"li", a, c, false}
}

func Main(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"main", a, c, false}
}

func Ol(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"ol", a, c, false}
}

func P(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"p", a, c, false}
}

func Pre(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"pre", a, c, false}
}

func Ul(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"ul", a, c, false}

}

// Inline Text Semantics
// TODO abbr, bdi, bdo, data, dfn, kbd, mark, q, rp, rt, rtc, ruby,
//      time, var, wbr

func A(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"a", a, c, false}
}

func B(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"b", a, c, false}
}

func Br(a string) *HtmlTree {
	return &HtmlTree{"br", a, []interface{}{}, true}
}

func Cite(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"cite", a, c, false}
}

func Code(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"code", a, c, false}
}

func Em(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"em", a, c, false}
}

func I(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"i", a, c, false}
}

func S(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"s", a, c, false}
}

func Samp(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"samp", a, c, false}
}

func Small(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"small", a, c, false}
}

func Span(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"span", a, c, false}
}

func Strong(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"strong", a, c, false}
}

func Sub(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"sub", a, c, false}
}

func Sup(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"sup", a, c, false}
}

func U(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"u", a, c, false}

}

// Image and Multimedia

func Area(a string) *HtmlTree {
	return &HtmlTree{"area", a, []interface{}{}, true}
}

func Audio(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"audio", a, c, false}
}

func Img(a string) *HtmlTree {
	return &HtmlTree{"img", a, []interface{}{}, true}
}

func Map(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"map", a, c, false}
}

func Track(a string) *HtmlTree {
	return &HtmlTree{"track", a, []interface{}{}, true}
}

func Video(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"video", a, c, false}

}

// Embedded interface{}

func Embed(a string) *HtmlTree {
	return &HtmlTree{"embed", a, []interface{}{}, true}
}

func Object(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"object", a, c, false}
}

func Param(a string) *HtmlTree {
	return &HtmlTree{"param", a, []interface{}{}, true}
}

func Source(a string) *HtmlTree {
	return &HtmlTree{"source", a, []interface{}{}, true}

}

// Scripting

func Canvas(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"canvas", a, c, false}
}

func Noscript(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"noscript", a, c, false}
}

func Script(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"script", a, c, false}

}

// Demarcating Edits
// TODO del, ins

// Table interface{}
// TODO colgroup (maybe. It's poorly supported.)

func Caption(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"caption", a, c, false}
}

func Col(a string) *HtmlTree {
	return &HtmlTree{"col", a, []interface{}{}, true}
}

func Table(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"table", a, c, false}
}

func Tbody(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"tbody", a, c, false}
}

func Td(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"td", a, c, false}
}

func Tfoot(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"tfoot", a, c, false}
}

func Th(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"th", a, c, false}
}

func Thead(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"thead", a, c, false}
}

func Tr(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"tr", a, c, false}

}

// Forms

func Button(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"button", a, c, false}
}

func Datalist(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"datalist", a, c, false}
}

func Fieldset(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"fieldset", a, c, false}
}

func Form(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"form", a, c, false}
}

func Input(a string) *HtmlTree {
	return &HtmlTree{"input", a, []interface{}{}, true}
}

func Label(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"label", a, c, false}
}

func Legend(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"legend", a, c, false}
}

func Meter(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"meter", a, c, false}
}

func Optgroup(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"optgroup", a, c, false}
}

func Option(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"option", a, c, false}
}

func Output(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"output", a, c, false}
}

func Progress(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"progress", a, c, false}
}

func Select(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"select", a, c, false}
}

func Textarea(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"textarea", a, c, false}

}

// Interactive Elememts (Experimental. Omitted for now.)

// Web Components (Experimental. Omitted for now.)
