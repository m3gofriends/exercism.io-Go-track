// Package bob classifies the message.
package bob

import "strings"

// Hey classifies the message.
func Hey(remark string) string {

	alphabetCount, lowerCount, upperCount, trimRemark := 0, 0, 0, strings.TrimSpace(remark)

	for _, letter := range trimRemark {
		if (letter >= 'a' && letter <= 'z') || (letter >= 'A' && letter <= 'Z') {
			alphabetCount++
			if letter >= 'a' && letter <= 'z' {
				lowerCount++
			} else {
				upperCount++
			}
		}
	}

	if len(trimRemark) == 0 {
		return "Fine. Be that way!"
	}

	switch alphabetCount {
	case 0:
		if trimRemark[len(trimRemark)-1] == '?' {
			return "Sure."
		}
	case upperCount:
		if trimRemark[len(trimRemark)-1] == '?' {
			return "Calm down, I know what I'm doing!"
		}
		return "Whoa, chill out!"
	default:
		if trimRemark[len(trimRemark)-1] == '?' {
			return "Sure."
		}
	}
	return "Whatever."
}
