package gohtx

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

// Null tag is a special case and does not correspond to any valid HTML tag.
// When rendered it returns its content with no enclosing tag, e.g.
// Null(Br(“),Br(“)) --> <br><br> . It is defined in goht to support injecting
// a list of elements into an existing tag as JavaScript innerHTML content.
func Null(c ...interface{}) *HtmlTree {
	return &HtmlTree{"null", ``, c, false}
}

// Document Metadata
// TODO: base

// Head, when rendered, returns a <head> element with the
// given attributes and content.
func Head(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"head", a, c, false}
}

// Body, when rendered, returns a <body> element with the given attributes and content.
func Body(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"body", a, c, false}
}

// Link, when rendered, returns a <link> element with the given attributes and content.
func Link(a string) *HtmlTree {
	return &HtmlTree{"link", a, []interface{}{}, true}
}

// Meta, when rendered, returns a <meta> element with the given attributes.
func Meta(a string) *HtmlTree {
	return &HtmlTree{"meta", a, []interface{}{}, true}
}

// Title, when rendered, returns a <title> element with the given attributes and content.
func Title(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"title", a, c, false}
}

// Style is a special case in the sense that the only valid content is one or
// more strings of CSS. At this time there's no check to complain about other
// content.
func Style(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"style", a, c, false}
}

// interface{} Sectioning
// TODO hgroup

// Address, when rendered, returns a <address> element with the given attributes and content.
func Address(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"address", a, c, false}
}

// Article, when rendered, returns an <article> element with the given attributes and content.
func Article(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"article", a, c, false}
}

// Aside, when rendered, returns an <aside> element with the given attributes and content.
func Aside(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"aside", a, c, false}
}

// Footer, when rendered, returns a <footer> element with the given attributes and content.
func Footer(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"footer", a, c, false}
}

// Header, when rendered, returns a <header> element with the given attributes and content.
func Header(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"header", a, c, false}
}

// H1, when rendered, returns a <h1> element with the given attributes and content.
func H1(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h1", a, c, false}
}

// H2, when rendered, returns a <h2> element with the given attributes and content.
func H2(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h2", a, c, false}
}

// H3, when rendered, returns a <h3> element with the given attributes and content.
func H3(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h3", a, c, false}
}

// H4, when rendered, returns a <h4> element with the given attributes and content.
func H4(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h4", a, c, false}
}

// H5, when rendered, returns a <h5> element with the given attributes and content.
func H5(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h5", a, c, false}
}

// H6, when rendered, returns a <h6> element with the given attributes and content.
func H6(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"h6", a, c, false}
}

// Nav, when rendered, returns a <nav> element with the given attributes and content.
func Nav(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"nav", a, c, false}
}

// Section, when rendered, returns a <section> element with the given attributes and content.
func Section(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"section", a, c, false}

}

// Details, when rendered, returns a <details> element with the given attributes and content.
func Details(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"details", a, c, false}
}

// Summary, when rendered, returns a <summary> element with the given attributes and content.
func Summary(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"summary", a, c, false}
}

// Text interface{}

// Blockquote, when rendered, returns a <blockquote> element with the given attributes and content.
func Blockquote(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"blockquote", a, c, false}
}

// Dd, when rendered, returns a <dd> element with the given attributes and content.
func Dd(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"dd", a, c, false}
}

// Div, when rendered, returns a <div> element with the given attributes and content.
func Div(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"div", a, c, false}
}

// Dl, when rendered, returns a <dl> element with the given attributes and content.
func Dl(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"dl", a, c, false}
}

// Dt, when rendered, returns a <dt> element with the given attributes and content.
func Dt(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"dt", a, c, false}
}

// Figcaption, when rendered, returns a <figcaption> element with the given attributes and content.
func Figcaption(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"figcaption", a, c, false}
}

// Figure, when rendered, returns a <figure> element with the given	 attributes and content.
func Figure(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"figure", a, c, false}
}

// Html, when rendered, returns an <hr> element with the given attributes.
func Hr(a string) *HtmlTree {
	return &HtmlTree{"hr", a, []interface{}{}, true}
}

// Li, when rendered, returns an <li> element with the given attributes and content.
func Li(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"li", a, c, false}
}

// Main, when rendered, returns an <main> element with the given attributes and content.
func Main(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"main", a, c, false}
}

// Ol, when rendered, returns an <ol> element with the given attributes and content.
func Ol(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"ol", a, c, false}
}

// P, when rendered, returns an <p> element with the given attributes and content.
func P(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"p", a, c, false}
}

// Pre, when rendered, returns a <pre> element with the given attributes and content.
func Pre(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"pre", a, c, false}
}

// Ul, when rendered, returns a <ul> element with the given attributes and content.
func Ul(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"ul", a, c, false}

}

// Inline Text Semantics
// TODO abbr, bdi, bdo, data, dfn, kbd, mark, q, rp, rt, rtc, ruby,
//      time, var, wbr

// A, when rendered, returns a <a> element with the given attributes and content.
func A(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"a", a, c, false}
}

// B, when rendered, returns a <b> element with the given attributes and content.
func B(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"b", a, c, false}
}

// Br, when rendered, returns a <br> element with the given attributes and content.
func Br(a string) *HtmlTree {
	return &HtmlTree{"br", a, []interface{}{}, true}
}

// Cite, when rendered, returns a <cite> element with the given attributes and content.
func Cite(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"cite", a, c, false}
}

// Code, when rendered, returns a <code> element with the given attributes and content.
func Code(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"code", a, c, false}
}

// Em, when rendered, returns a <em> element with the given attributes and content.
func Em(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"em", a, c, false}
}

// I, when rendered, returns an <i> element with the given attributes and content.
func I(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"i", a, c, false}
}

// S, when rendered, returns an <s> element with the given attributes and content.
func S(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"s", a, c, false}
}

// Samp, when rendered, returns a <samp> element with the given attributes and content.
func Samp(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"samp", a, c, false}
}

// Small, when rendered, returns a <small> element with the given attributes and content.
func Small(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"small", a, c, false}
}

// Span, when rendered, returns a <span> element with the given attributes and content.
func Span(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"span", a, c, false}
}

// Strong, when rendered, returns a <strong> element with the given attributes and content.
func Strong(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"strong", a, c, false}
}

// Sub, when rendered, returns a <sub> element with the given attributes and content.
func Sub(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"sub", a, c, false}
}

// Sup, when rendered, returns a <sup> element with the given attributes and content.
func Sup(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"sup", a, c, false}
}

// U, when rendered, returns a <u> element with the given attributes and content.
func U(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"u", a, c, false}

}

// Image and Multimedia

// Area, when rendered, returns an <area> element with the given attributes.
func Area(a string) *HtmlTree {
	return &HtmlTree{"area", a, []interface{}{}, true}
}

// Audio, when rendered, returns a <audio> element with the given attributes and content.
func Audio(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"audio", a, c, false}
}

// Img, when rendered, returns a <img> element with the given attributes.
func Img(a string) *HtmlTree {
	return &HtmlTree{"img", a, []interface{}{}, true}
}

// Map, when rendered, returns a <map> element with the given attributes and content.
func Map(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"map", a, c, false}
}

// Track, when rendered, returns a <track> element with the given attributes.
func Track(a string) *HtmlTree {
	return &HtmlTree{"track", a, []interface{}{}, true}
}

// Video, when rendered, returns a <video> element with the given attributes and content.
func Video(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"video", a, c, false}

}

// Embedded interface{}

// Embed, when rendered, returns an <embed> element with the given attributes.
func Embed(a string) *HtmlTree {
	return &HtmlTree{"embed", a, []interface{}{}, true}
}

// Object, when rendered, returns a <object> element with the given attributes and content.
func Object(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"object", a, c, false}
}

// Param, when rendered, returns a <param> element with the given attributes.
func Param(a string) *HtmlTree {
	return &HtmlTree{"param", a, []interface{}{}, true}
}

// Source, when rendered, returns a <source> element with the given attributes.
func Source(a string) *HtmlTree {
	return &HtmlTree{"source", a, []interface{}{}, true}

}

// Scripting

// Canvas, when rendered, returns a <canvas> element with the given attributes and content.
func Canvas(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"canvas", a, c, false}
}

// Noscript, when rendered, returns a <noscript> element with the given attributes and content.
func Noscript(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"noscript", a, c, false}
}

// Script, when rendered, returns a <script> element with the given attributes and content.
func Script(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"script", a, c, false}

}

// Demarcating Edits
// TODO del, ins

// Table interface{}
// TODO colgroup (maybe. It's poorly supported.)

// Caption, when rendered, returns a <caption> element with the given attributes and content.
func Caption(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"caption", a, c, false}
}

// Col, when rendered, returns a <col> element with the given attributes and content.
func Col(a string) *HtmlTree {
	return &HtmlTree{"col", a, []interface{}{}, true}
}

// Table, when rendered, returns a <table> element with the given attributes and content.
func Table(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"table", a, c, false}
}

// Tbody, when rendered, returns a <tbody> element with the given attributes and content.
func Tbody(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"tbody", a, c, false}
}

// Td, when rendered, returns a <td> element with the given attributes and content.
func Td(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"td", a, c, false}
}

// Tfoot, when rendered, returns a <tfoot> element with the given attributes and content.
func Tfoot(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"tfoot", a, c, false}
}

// Th, when rendered, returns a <th> element with the given attributes and content.
func Th(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"th", a, c, false}
}

// Thead, when rendered , returns a <thead> element with the given attributes and content.
func Thead(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"thead", a, c, false}
}

// Tr, when rendered, returns a <tr> element with the given attributes and content.
func Tr(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"tr", a, c, false}

}

// Forms

// Button, when rendered, returns a <button> element with the given attributes and content.
func Button(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"button", a, c, false}
}

// Datalist, when rendered, returns a <datalist> element with the given attributes and content.
func Datalist(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"datalist", a, c, false}
}

// Fieldset, when rendered, returns a <fieldset> element with the given attributes and content.
func Fieldset(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"fieldset", a, c, false}
}

// Form, when rendered, returns a <form> element with the given attributes and content.
func Form(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"form", a, c, false}
}

// Input, when rendered, returns a <input> element with the given attributes and content.
func Input(a string) *HtmlTree {
	return &HtmlTree{"input", a, []interface{}{}, true}
}

// Label, when rendered, returns a <label> element with the given attributes and content.
func Label(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"label", a, c, false}
}

// Legend, when rendered, returns a <legend> element with the given attributes and content.
func Legend(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"legend", a, c, false}
}

// Meter, when rendered, returns a <meter> element with the given attributes and content.
func Meter(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"meter", a, c, false}
}

// Optgroup, when rendered, returns a <optgroup> element with the given attributes and content.
func Optgroup(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"optgroup", a, c, false}
}

// Option, when rendered, returns a <option> element with the given attributes and content.
func Option(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"option", a, c, false}
}

// Output, when rendered, returns a <output> element with the given attributes and content.
func Output(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"output", a, c, false}
}

// Progress, when rendered, returns a <progress> element with the given attributes and content.
func Progress(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"progress", a, c, false}
}

// Select, when rendered, returns a <select> element with the given attributes and content.
func Select(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"select", a, c, false}
}

// Textarea, when rendered, returns a <textarea> element with the given attributes and content.
func Textarea(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"textarea", a, c, false}

}

// As of 2021, <dialog> is supported by Chrome, Edge & Opera but
// not by FireFox and Safari.
// Dialog, when rendered, returns a <dialog> element with the given attributes and content.
func Dialog(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"dialog", a, c, false}
}

// Interactive Elememts (Experimental. Omitted for now.)

// Web Components (Experimental. Omitted for now.)
