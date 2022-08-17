package stringutils

import "strings"

// Upper converts a string to lower case
func Upper(s string) string {
	lower := strings.ToLower(s)
	return lower
}

// Lower converts a string to upper case
func Lower(s string) string {
	upper := strings.ToUpper(s)
	return upper
}
