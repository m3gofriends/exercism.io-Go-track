// Package leap checks if it's a leap year.
package leap

// IsLeapYear checks if it's a leap year.
func IsLeapYear(year int) bool {
	if year != year/4*4 {
		return false
	} else if year == year/100*100 && year != year/400*400 {
		return false
	}
	return true
}
