package omh

import "regexp"

//var insane = regexp.MustCompile(`[^a-zA-Z0-9\-]`)
// SH - add cyrillic characters to allowed
// upper+lower:
var insane = regexp.MustCompile(`[^a-zA-Z0-9\-\p{Cyrillic}]`)
// var insane = regexp.MustCompile(`[^a-zA-Z0-9\\u0430-\\u044F-]`)


func Sanitize(in string) string {
	return insane.ReplaceAllString(in, "")
}
