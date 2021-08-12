// Package isogram checks if string has duplicate letters.
package isogram

import "strings"

var letter = make(map[rune]bool)

// IsIsogram checks if string has duplicate letters.
func IsIsogram(s string) bool {

	letter = make(map[rune]bool)

	for _, value := range strings.ToUpper(s) {
		if value == '-' || value == ' ' {
			continue
		}
		if letter[value] {
			return false
		}
		letter[value] = true
	}
	return true
}
