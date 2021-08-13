// Package encode implements run-length encoding and decoding.
package encode

import "strconv"

// RunLengthEncode implements run-length encoding.
func RunLengthEncode(s string) (output string) {
	if len(s) == 0 {
		return ""
	}

	nowLetter, counter := s[0], 1
	for i := 1; i < len(s); i++ {
		if nowLetter == s[i] {
			counter++
		} else {
			if counter == 1 {
				output += string(nowLetter)
			} else {
				output += strconv.Itoa(counter) + string(nowLetter)
			}
			nowLetter = s[i]
			counter = 1
		}
	}
	if counter == 1 {
		output += string(nowLetter)
	} else {
		output += strconv.Itoa(counter) + string(nowLetter)
	}
	return output
}

// RunLengthDecode implements run-length decoding.
func RunLengthDecode(s string) (output string) {
	digits := ""

	for i := 0; i < len(s); i++ {
		if s[i] > 47 && s[i] < 58 { // 0 to 9
			digits += string(s[i])
		} else {
			if digits == "" {
				output += string(s[i])
			} else {
				digit, _ := strconv.Atoi(digits)
				for times := 0; times < digit; times++ {
					output += string(s[i])
				}
				digits = ""
			}
		}
	}
	return output
}
