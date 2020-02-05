/* package accumulate Applies a given function to create a new slice of strings. Has TDT. */
package accumulate

//Accumulate Outputs a new slice that looks like the function applied to the old slice.
func Accumulate(in []string, f func(string) string) (out []string) {
	out = make([]string, len(in))
	for i, s := range in {
		out[i] = f(s)
	}
	return
}
