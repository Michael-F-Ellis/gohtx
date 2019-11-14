# goht
### A minimalist package for creating web content in Go.

## Features
* Easy to learn if you already know html.
* Dynamic. Generate html/css/js at build time or run time, server side or client side (using GopherJS or WebAssembly).
* Over 80 common tags pre-defined. Adding new ones is trivially simple. 
* Composable. Reduce repetition by writing functions that return nested html fragments.
* Fast. About 100ns per tag on typical hardware.

## Usage
```import . "github.com/Michael-F-Ellis/goht"```

Dot imports are generally discouraged as bad style but they make sense for this package. It saves you having to put `goht.` in front of every tag function. 

### Hello, World
```
var page := Html("",Head(""),Body("", "Hello, World"))
var b bytes.Buffer
err := Render(page, &b, 0)
```

produces

```
<html>
  <head>
  </head>
  <body>Hello, World
  </body>
</html>
```

See [Example 1](https://goplay.space/#eP-DfcNaJxh) for a complete listing you can edit and run. 

For a larger usage example, see `webpage.go` in the [repository](https://github.com/Michael-F-Ellis/infinite-etudes) for my online ear-training app [Infinite Etudes](https://etudes.ellisandgrant.com).

### Composing
Our "Hello, World" example

can also be written
```
head := Head("")
body := Body("", "Hello, World")
page := Html("", head, body)
```

or, possibly (see [Example 2](https://goplay.space/#l2PWufxkgLV))

```
func mkPage(c ...interface{}) *HtmlTree {
    body := Body("", c...)
    return Html("", Head(""), body)
}
page1 := mkPage("Hello, World")
page2 := mkPage(H1("", "Mars"), "Hello, Other World")
```

The point is that the full power and flexibility of Go is available to factor repetitive content. Any function that returns a valid `*HtmlTree` may be called to create content arguments for any tag function.

## Under the hood
### Tags

The definition for all (non-empty) html tags is the same:

```
func Tagname(a string, c ...interface{}) *HtmlTree {
	return &HtmlTree{"tagname", a, c, false}
}
```

The func name is simply the html tag name with an initial cap.

The first arg, `a`, is a string of attributes written exactly as you would in html, e.g. `id=42 class="foo"`. In `goht` it's helpful to enclose the attribute string in back-quotes to allow use of both single and double quotes when specifying attributes.

The second arg, `c`, is the content of the tag.  It's a variadic argument meaning that you may supply as many arguments as needed to define the inner html of the tag.  The type of `c` is `interface{}`. Only two concrete types are supported: `string` or `*HtmlTree`. The latter is the return value type of every tag function.
  
Empty tags, like `<br>` can't contain other elements. In `goht` these tag have only the `a` argument, e.g 
```
Br(a string) *HtmlTree { ... }
```
See `tags.go` for the complete list of tags defined in `goht`.

### The HtmlTree struct

`Goht` represents nested html elements with a recursive struct:

```
// HtmlTree represents a tree of html content.
type HtmlTree struct {
	T     string        // html tagname, e.g. 'head'
	A     string        // zero or more html attributes, e.g 'id=1 class="foo"'
	C     []interface{} // a slice of content whose elements may be strings or *HtmlTree
	empty bool          // set to true for empty tags like <br>
}
```
You'll mostly use HtmlTree as a return type when composing grouped tag functions. The struct members are exported should you need to access them from within some clever on-the-fly html generation, but as Rob Pike says "Clear is better than clever. Don't be clever."



### Rendering
A single func is exposed to handle rendering html text from HtmlTree structs.
```
func Render(h *HtmlTree, b *bytes.Buffer, nindent int) (err error)
```

Render walks through `HtmlTree h` and writes html text to byte buffer `b`. The
`nindent` argument specifies whether and how much to indent the output where `-1`
means render all on one line and 0 means indent each inner tag by `2` spaces
relative to its parent. Render returns an `error` when it encounters invalid
content.

## CSS
Goht provides the Style tag function. It's up to you to make sure the contents are valid CSS.
```
Style("", `
    body {
	  margin: 0;
	  height: 100%;
	  overflow: auto;
	  background-color: #DDA;
	  }
    h1 {font-size: 300%; margin-bottom: 1vh}
    `)
```
You can also import external stylesheets with the Link tag function. For example
```
Link(`rel="stylesheet" href="https://www.w3schools.com/w3css/4/w3.css"`)
```

## JavaScript
Use the Script tag function to import or contain JavaScript code. 

```
Script("src=/midijs/midi.js charset=UTF-8")

Script("",
		`
		// chores at start-up
		function start() {
		  // Chrome and other browsers now disallow AudioContext until
		  // after a user action.
		  document.body.addEventListener("click", MIDIjs.resumeAudioContext);
		}
		// Run start when the doc is fully loaded.
		document.addEventListener("DOMContentLoaded", start);
	`)
```
## Checking attributes
Goht includes a limited facility for checking tag attributes. It's mainly useful for unit tests.
```
func (e *HtmlTree) CheckAttributes(perrs *[]AttributeErrors)
```
CheckAttributes walks through an HtmlTree and checks each tag to verify that
the attribute names associated with the tag are valid for that tag. It returns a slice of AttributeErrors. The slice will be empty if no errors were found.

```
type AttributeErrors struct {
	tag   string  // an html tag
	attrs string  // a string of zero or more attributes
	errs  []error // errors found in attrs
}
```

See [Example 3](https://goplay.space/#UYp7qPBfXq7) for usage details

## Alternatives
Goht is designed with a "simplest thing that could possibly work" philosophy. Here are some more ambitious alternatives.

* [vugu](https://www.vugu.org/)
* [vecty](https://github.com/gopherjs/vecty)
* [html/template](https://golang.org/pkg/html/template/)