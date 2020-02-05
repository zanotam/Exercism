/*
Package etl Corresponds to the T in ETL and Transforms old scrabble data into a new form.
*/
package etl

import "strings"

//Transform Transforms the original style data into the new requirement style data.
func Transform(og map[int][]string) (transformed map[string]int) {
	transformed = make(map[string]int)
	for key, values := range og {
		for _, value := range values {
			transformed[strings.ToLower(value)] = key
		}
	}
	return
}
