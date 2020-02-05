// Package proverb Generates a 'proverb' based upon a list of rhyming input.
// Also includes tests for making sure the whole package works right.
package proverb

// Proverb Generates a 'proverb' based upon a list of 'rhyming' input in the form of a string slice.
func Proverb(rhyme []string) []string {
	var poem []string = []string{}
	//the repetitive chorus-like verses that make up the majority of the proverb
	for i := 1; len(rhyme) > i; i++ {
		verse := "For want of a "
		verse += rhyme[i-1]
		verse += " the "
		verse += rhyme[i]
		verse += " was lost."
		poem = append(poem, verse)
	}
	//the final verse which is different than the repetitive chorus-like verses... the.... message of the story, I guess
	if len(rhyme) > 0 {
		verse := "And all for the want of a "
		verse += rhyme[0]
		verse += "."
		poem = append(poem, verse)
	}
	return poem
}
