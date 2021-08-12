// Package isogram checks if string has duplicate letters.
package isogram

import "unicode"

var letter = make(map[rune]bool)

// IsIsogram checks if string has duplicate letters.
func IsIsogram(s string) bool {
	letter = make(map[rune]bool)

	for _, value := range s {
		if !unicode.IsLetter(value) {
			continue
		}

		value = unicode.ToLower(value)
		if _, ok := letter[value]; ok {
			return false
		}
		letter[value] = true
	}
	return true
}
