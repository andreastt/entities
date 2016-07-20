// Convert special HTML characters to entities, and vice versa.

package entities

import (
	"html"
	"strings"
)

var entityEscaper = strings.NewReplacer(
	"&", "&amp;",
	"'", "&apos;",
	"<", "&lt;",
	">", "&gt;",
	"\"", "&quot;",
)

// Escape special entity characters in HTML documents, so that characters
// like "<" become "&lt;".
//
// It escapes a wider range of entities than the standard library html
// package.  Unescape(Escape(s)) == s is always true, but the converse
// cannot be guaranteed.
func Escape(s string) string {
	return entityEscaper.Replace(s)
}

// Unescape entities like "&lt;" to become "<". It unescapes a larger
// range of entities than Escape escapes.
//
// This is currently an alias for html.UnescapeString.  See its
// documentation.
func Unescape(s string) string {
	return html.UnescapeString(s)
}
