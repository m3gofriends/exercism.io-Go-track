// Package pangram determines if a sentence is a pangram.
package pangram

// IsPangram determines if a sentence is a pangram.
func IsPangram(s string) bool {
	var letterFlags uint32 = 0

	for i := 0; i < len(s); i++ {
		letter := s[i]
		if letter > 64 && letter < 91 { // A to Z
			letter += 32
		}

		if letter > 96 && letter < 123 { // a to z
			letterFlags |= (1 << (letter - 'a'))
		}
	}
	return letterFlags == 67108863 // 67108863 equals 111111111111111111111111 in base-2
}
