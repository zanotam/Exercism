/*
Package hamming Calculates the hamming distance between two strings.
Strings are meant to represent DNA code, but this package generalizes the calculation of the Hamming distance
so arbitrary strings can be used.
*/
package hamming

import "errors"

//Distance Calculates the Hamming distance of the strings or returns an error if the strings are of different length.
func Distance(a, b string) (metric int, err error) {
	//error handling
	a2 := []rune(a)
	b2 := []rune(b)
	if len(a2) != len(b2) {
		err = errors.New("the hamming distance cannot be calculated on strands with different base length")
		return
	}
	//actual hamming distance calculation
	for index := range a2 {
		if a2[index] != b2[index] {
			metric++
		}
	}
	return
}
