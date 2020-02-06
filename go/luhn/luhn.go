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
	strippedDigits := strings.ReplaceAll(digits, " ", "")
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
	for i, digit := range digits {
		//convert character into a number
		var sliceDigit = make([]rune, 1)
		sliceDigit[0] = digit
		strDigit := string(sliceDigit)
		numDigit, errNumberConversion := strconv.Atoi(strDigit) //I think plain 'err' is too ambiguous. Plus it's already declared!
		if errNumberConversion != nil {
			return errNumberConversion, 0
		}
		//check if it's legit
		shouldDouble := (((len(digits) - 1 - i) % 2) == 1) && (i != (len(digits) - 1))
		if shouldDouble {
			numDigit *= 2
			if numDigit > 9 {
				numDigit -= 9
			}
		}
		summed += numDigit
	}
	return nil, summed
}
