/*
Package pangram Determines if a string is a pangram. Includes TDT.
*/
package pangram

import "unicode"

const a = 'a'
const z = 'z'

//IsPangram Determines if a given string is a pangram.
func IsPangram(in string) bool {
	letterSet := make(map[rune]bool)
	inLetters := []rune(in)
	for _, letter := range inLetters {
		lower := unicode.ToLower(letter)
		letterSet[lower] = true
	}
	success := true
	for i := a; i <= z; i++ {
		if !letterSet[i] {
			success = false
		}
	}
	return success
}
