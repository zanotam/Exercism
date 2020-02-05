/*
package romannumerals converts integers into their roman numeral representation and includes TDT for this functionality.
*/
package romannumerals

import (
	"errors"
	"strconv"
)

//ToRomanNumeral Converts an integer into a roman numeral representation if in the range of [1,3000].
func ToRomanNumeral(arabic int) (roman string, err error) {
	if (arabic > 3000) || (arabic < 1) {
		err = errors.New("No legitimate corresponding roman numeral.")
		return
	}
	digits := []rune(strconv.Itoa(arabic))
	if len(digits) > 3 {
		str := string(digits[0])
		switch str {
		case "1":
			roman += "M"
		case "2":
			roman += "MM"
		case "3":
			roman += "MMM"
		}
		digits = digits[1:]
	}
	if len(digits) > 2 {
		str := string(digits[0])
		switch str {
		case "9":
			roman += "CM"
		case "8":
			roman += "DCCC"
		case "7":
			roman += "DCC"
		case "6":
			roman += "DC"
		case "5":
			roman += "D"
		case "4":
			roman += "CD"
		case "3":
			roman += "CCC"
		case "2":
			roman += "CC"
		case "1":
			roman += "C"
		}
		digits = digits[1:]
	}
	if len(digits) > 1 {
		str := string(digits[0])
		switch str {
		case "9":
			roman += "XC"
		case "8":
			roman += "LXXX"
		case "7":
			roman += "LXX"
		case "6":
			roman += "LX"
		case "5":
			roman += "L"
		case "4":
			roman += "XL"
		case "3":
			roman += "XXX"
		case "2":
			roman += "XX"
		case "1":
			roman += "X"
		}
		digits = digits[1:]
	}
	str := string(digits)
	switch str {
	case "9":
		roman += "IX"
	case "8":
		roman += "VIII"
	case "7":
		roman += "VII"
	case "6":
		roman += "VI"
	case "5":
		roman += "V"
	case "4":
		roman += "IV"
	case "3":
		roman += "III"
	case "2":
		roman += "II"
	case "1":
		roman += "I"
	}
	return
}
