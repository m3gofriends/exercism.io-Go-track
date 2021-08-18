// Package isbn checks if the provided string is a valid ISBN-10.
package isbn

// IsValidISBN checks if the provided string is a valid ISBN-10.
func IsValidISBN(s string) bool {
	counter, sum := 0, 0

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			sum += int(s[i]-'0') * (10 - counter)
			counter++
		case 'X':
			if i != len(s)-1 {
				return false
			}
			sum += 10
			counter++
		case '-':
		default:
			return false
		}
	}

	if counter != 10 {
		return false
	}

	return sum%11 == 0
}
