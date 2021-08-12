// Package pangram determines if a sentence is a pangram.
package pangram

import "unicode"

// IsPangram determines if a sentence is a pangram.
func IsPangram(s string) bool {

	letter := make(map[rune]bool)

	for _, value := range s {
		value = unicode.ToLower(value)
		if !unicode.IsLetter(value) && !letter[value] {
			continue
		}
		letter[value] = true
	}
	return len(letter) == 26
}
