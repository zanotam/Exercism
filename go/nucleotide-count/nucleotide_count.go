/* package dna is used to create histograms for DNA strands including TDT. */
package dna

import "errors"

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	h = make(map[rune]int)
	h['A'] = 0
	h['C'] = 0
	h['G'] = 0
	h['T'] = 0
	for _, char := range d {
		var s = rune(char)
		if (s == 'A') || (s == 'C') || (s == 'G') || (s == 'T') {
			h[s] += 1
		} else {
			return nil, errors.New("non-DNA strand cannot have a histogram created")
		}
	}
	return h, nil
}
