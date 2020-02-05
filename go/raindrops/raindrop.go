/*
Package raindrops Converts a number to a string, the contents of which depend on the number's factors.
Has testing for this feature.
*/
package raindrops

import "strconv"

//soundMap maps the factors to the appropriate string "sound".
var soundMap = []struct {
	num   int
	sound string
}{
	{3, "Pling"},
	{5, "Plang"},
	{7, "Plong"},
}

//Convert Converts a number to a string, the contents of which depend on the number's factors.
func Convert(num int) (s string) {
	for _, x := range soundMap {
		if num%x.num == 0 {
			s += x.sound
		}
	}
	if s == "" {
		s = strconv.Itoa(num) //relies on the fact that strings are equivalent to rune slices
	}
	return
}
