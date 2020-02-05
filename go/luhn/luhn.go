/*
Package luhn checks if a string passes the Luhn test. Includes TDT.
*/
package luhn

import (
	"strconv"
	"strings"
)

//Valid Checks if a string passes the Luhn test
func Valid(digits string) (isValid bool) {
	strippedDigits := strings.Replace(digits, " ", "", -1)
	err, summedDigits := luhnDoubleAndSum(strippedDigits)
	if err == nil {
		isValid = (summedDigits%10) == 0 && (len(strippedDigits) > 1)
		return isValid
	}
	return false
}

//switched to error first because I like that convention
func luhnDoubleAndSum(digits string) (err error, summed int) {
	//var doubled string;
	for i, digit := range []rune(digits) {
		//convert character into a number
		var array = make([]rune, 1)
		array[0] = digit
		str := string(array)
		number, numberCheck := strconv.Atoi(str)
		//check if it's legit
		if (((len(digits) - 1 - i) % 2) == 1) && (i != (len(digits) - 1)) {
			number *= 2
			if number > 9 {
				number -= 9
			}
		}
		//convert back to a string
		//doubled = strconv.Itoa(number) + doubled
		if numberCheck != nil {
			return numberCheck, 0
		}
		summed += number
	}
	return nil, summed
}
