// Package luhn checks whether a given string is valid according to the Luhn formula.
package luhn

// Valid checks whether a given string is valid according to the Luhn formula.
func Valid(s string) bool {
	counter, sum := 0, 0

	for i := len(s) - 1; i >= 0; i-- {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			if counter%2 == 0 {
				sum += int(s[i] - '0')
			} else {
				if s[i] > '4' {
					sum += int(s[i]-'0')*2 - 9
				} else {
					sum += int(s[i]-'0') * 2
				}
			}
			counter++
		case ' ':
		default:
			return false
		}
	}

	if counter < 2 {
		return false
	}

	return sum%10 == 0
}
