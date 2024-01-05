package parser

import (
	"regexp"
	"strings"
)

// alphanumeric will remove from string all non-alphanumeric characters
func alphanumeric(str string) string {

	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9 ]+`)

	return nonAlphanumericRegex.ReplaceAllString(str, "")
}

// noWhiteSpaces will remove all white spaces from a string
func noWhiteSpaces(str string) string {

	return strings.Join(strings.Fields(str), "")
}
