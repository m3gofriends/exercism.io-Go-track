// Package leap checks if it's a leap year.
package leap

// IsLeapYear checks if it's a leap year.
func IsLeapYear(year int) bool {
	if year%100 == 0 {
		return year%400 == 0
	}
	return year%4 == 0
}
