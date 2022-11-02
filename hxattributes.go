package gohtx

// HxAttrs is a slice of string containing the names of all valid htmx attributes.
// The names ad descriptions are taken from https://htmx.org/reference/
var HxAttrs []string = []string{
	// Core attributes
	"hx-boost",      // add or remove progressive enhancement for links and forms
	"hx-get",        // issues a GET to the specified URL
	"hx-post",       // issues a POST to the specified URL
	"hx-push-url",   // pushes the URL into the browser location bar, creating a new history entry
	"hx-select",     // select content to swap in from a response
	"hx-select-oob", // select content to swap in from a response, out of band (somewhere other than the target)
	"hx-swap",       // controls how content is swapped in (outerHTML, beforeEnd, afterend, ...)
	"hx-swap-oob",   // marks content in a response to be out of band (should swap in somewhere other than the target)
	"hx-target",     // specifies the target element to be swapped
	"hx-trigger",    // specifies the event that triggers the request
	"hx-vals",       // adds values to the parameters to submit with the request (JSON-formatted)

	// Additional Attributes
	"hx-confirm",     // shows a confim() dialog before issuing a request
	"hx-delete",      // issues a DELETE to the specified URL
	"hx-disable",     // disables htmx processing for the given node and any children nodes
	"hx-disinherit",  // control and disable automatic attribute inheritance for child nodes
	"hx-encoding",    // changes the request encoding type
	"hx-ext",         // extensions to use for this element
	"hx-headers",     // adds to the headers that will be submitted with the request
	"hx-history-elt", // the element to snapshot and restore during history navigation
	"hx-include",     // include additional data in requests
	"hx-indicator",   // the element to put the htmx-request class on during the request
	"hx-params",      // filters the parameters that will be submitted with a request
	"hx-patch",       // issues a PATCH to the specified URL
	"hx-preserve",    // specifies elements to keep unchanged between requests
	"hx-prompt",      // shows a prompt() before submitting a request
	"hx-put",         // issues a PUT to the specified URL
	"hx-replace-url", // replace the URL in the browser location bar
	"hx-request",     // configures various aspects of the request
	"hx-sse",         // has been moved to an extension. Documentation for older versions
	"hx-sync",        // control how requests made be different elements are synchronized
	"hx-vars",        // adds values dynamically to the parameters to submit with the request (deprecated, please use hx-vals)
	"hx-ws",          // has been moved to an extension. Documentation for older versions
	"_",              // (HyperScript) the convenient if slightly cryptic alternative to "script"
	"script",         // (HyperScript) canonical name for a script attribute
}
