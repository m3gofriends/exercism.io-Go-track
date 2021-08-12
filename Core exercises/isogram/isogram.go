// Package isogram checks if string has duplicate letters.
package isogram

import "strings"

// IsIsogram checks if string has duplicate letters.
func IsIsogram(s string) bool {
	s = strings.ToLower(s)
	for _, value := range s {
		if (value >= 'a' && value <= 'z') && (strings.Count(s, string(value)) > 1) {
			return false
		}
	}
	return true
}
