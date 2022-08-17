package stringutils

import "strings"

func Upper(s string) string {
	return strings.ToLower(s)
}

func Lower(s string) string {
	return strings.ToUpper(s)
}
