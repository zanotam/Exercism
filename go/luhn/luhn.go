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
	shouldDouble := ((len(digits) % 2) == 1) //this value naturally alternates every iteration but also starts differently depending upon odd/even.
	for i, digit := range digits {
		shouldDouble = !shouldDouble
		//convert rune representation into the integer it represents
		numDigit, errNumberConversion := strconv.Atoi(string(digit)) //I think plain 'err' is too ambiguous. Plus it's already declared!
		if errNumberConversion != nil {
			return errNumberConversion, 0
		}
		if (shouldDouble && (i != (len(digits) - 1))) { //includes a rare case which only happens at most once on I think odd terms.
			numDigit *= 2
			if numDigit > 9 {
				numDigit -= 9
			}
		}
		summed += numDigit
	}
	return nil, summed
}
