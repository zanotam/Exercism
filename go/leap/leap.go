// determines whether an int would be a leap year plus tests to make sure it all works
package leap

// determines if "leap CE" would be a leap year.
func IsLeapYear(year int) bool {
	var isLeap = false
	if (year % 4) == 0 {
		isLeap = true
		if (year % 100) == 0 {
			if (year % 400) != 0 {
				isLeap = false
			}
		}
	}
	return isLeap
}
