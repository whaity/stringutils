package stringutils

import "strings"

// Upper converts a string to lower case
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Lower converts a string to upper case
func Lower(s string) string {
	return strings.ToLower(s)
}
