// Package raindrops converts a number into a string.
package raindrops

import "strconv"

// Convert converts a number into a string.
func Convert(number int) (result string) {
	switch number {
	case number / 3 * 3, number / 5 * 5, number / 7 * 7:
		if number == number/3*3 {
			result += "Pling"
		}
		if number == number/5*5 {
			result += "Plang"
		}
		if number == number/7*7 {
			result += "Plong"
		}
		return result
	default:
		return strconv.Itoa(number)
	}
}
