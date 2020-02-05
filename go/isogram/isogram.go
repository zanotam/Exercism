/*
Package isogram Determines whether a word is an isogram. Includes TDT.
*/
package isogram

import "unicode"

//IsIsogram Checks if a given string is an isogram (has no repeated letters).
func IsIsogram(word string) (pass bool) {
	pass = true
	letterSet := make(map[rune]bool)
	for _, char := range word {
		if letterSet[unicode.ToUpper(char)] && char != '-' && char != ' ' {
			pass = false
		}
		letterSet[unicode.ToUpper(char)] = true
	}
	return
}
