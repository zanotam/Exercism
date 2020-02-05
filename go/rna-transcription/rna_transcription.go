/* package strand turns a string representing a DNA strand into the RNA complement strand including TDT */
package strand

//complementMap maps the DNA letter to the appropriate RNA letter
var complementMap = map[string]string{
	"G": "C",
	"C": "G",
	"T": "A",
	"A": "U",
}

//ToRNA transcribes a DNA string into an appropriate RNA string.
func ToRNA(dna string) string {
	rna := ""
	for _, char := range dna {
		rna += complementMap[string(char)]
	}
	return rna
}
