// Package raindrops converts a number into a string.
package raindrops

import "strconv"

// Convert converts a number into a string.
func Convert(number int) (result string) {
	if number%3 == 0 {
		result += "Pling"
	}
	if number%5 == 0 {
		result += "Plang"
	}
	if number%7 == 0 {
		result += "Plong"
	}
	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}
