package main

import (
	"strings"
	"unicode"
)

/*
The purpose of this function will be to split the users input into "words" based on whitespace.

It should also lowercase the input and trim any leading or trailing whitespace. For example:
*/
func cleanInput(text string) []string {
	if text == "" {
		return []string{}
	}
	message := strings.ToLower(text)
	// only returns letters and spaces, removes everything else.
	message = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, message)
	return strings.Fields(message)
}
