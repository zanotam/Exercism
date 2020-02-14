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
	digits = strings.ReplaceAll(digits, " ", "")
	sum, err := luhnDoubleAndSum(digits)
	if err == nil {
		isValid = (sum%10) == 0 && (len(digits) > 1)
		return isValid
	}
	return false
}

func luhnDoubleAndSum(digits string) (sum int, err error) {
	//shouldDOuble naturally alternates every iteration but also starts differently depending upon odd/even.
	shouldDouble := len(digits)%2 == 1
	for i, digit := range digits {
		shouldDouble = !shouldDouble
		numDigit, err := strconv.Atoi(string(digit))
		if err != nil {
			return 0, err
		}
		//includes a rare case which only happens at most once on I think odd terms.
		if shouldDouble && (i != (len(digits) - 1)) {
			numDigit *= 2
			if numDigit > 9 {
				numDigit -= 9
			}
		}
		sum += numDigit
	}
	return sum, nil
}
