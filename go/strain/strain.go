/*
Package straing Implements 'keep' and 'discard' functions for custom types. Includes testing.
*/
package strain

//mandatory types for the tests
type Ints []int
type Lists [][]int
type Strings []string

//Ints.Keep Gives kept according to the predicate.
func (list Ints) Keep(f func(int) bool) Ints {
	if list != nil {

		var chosen Ints = make([]int, 0)
		for _, num := range list {
			if f(num) {
				chosen = append(chosen, num)
			}
		}
		return chosen
	}
	return nil
}

//Ints.Discard Gives discarded according to the predicate.
func (list Ints) Discard(f func(int) bool) Ints {
	if list != nil {
		var chosen Ints = make([]int, 0)
		for _, num := range list {
			if !f(num) {
				chosen = append(chosen, num)
			}
		}
		return chosen
	}
	return nil
}

//Lists.Keep Gives kept according to the predicate.
func (lists Lists) Keep(f func([]int) bool) Lists {
	if lists != nil {
		var chosen Lists = make([][]int, 0)
		for _, list := range lists {
			if f(list) {
				chosen = append(chosen, list)
			}
		}
		return chosen
	}
	return nil
}

//Strings.Keep Gives kept according to the predicate
func (list Strings) Keep(f func(string) bool) Strings {
	if list != nil {
		var chosen Strings = make([]string, 0)
		for _, str := range list {
			if f(str) {
				chosen = append(chosen, str)
			}
		}
		return chosen
	}
	return nil
}
