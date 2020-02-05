/*
Package reverse Reverses the order of runes in a string. Includes tests.
*/
package reverse

//Reverse Reverses the order of runes in a string.
func Reverse(str string) (reversed string) {
	var s, r []rune
	r = []rune(str)
	if len(r) < 2 {
		reversed = string(r)
	} else {
		for i := 1; i <= len(r); i++ {
			s = append(s, r[len(r)-i])
		}
		reversed = string(s)
	}
	return
}
