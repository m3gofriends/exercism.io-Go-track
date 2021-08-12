// Package isogram checks if string has duplicate letters.
package isogram

// IsIsogram checks if string has duplicate letters.
func IsIsogram(s string) bool {
	var letterFlags uint32 = 0

	for i := 0; i < len(s); i++ {
		letter := s[i]
		if letter > 96 && letter < 123 {
			letter -= 32
		}

		if letter > 64 && letter < 91 {
			if (letterFlags & (1 << (letter - 'A'))) != 0 {
				return false
			}
			letterFlags |= (1 << (letter - 'A'))
		}
	}
	return true
}
