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
	if alphabetCount == upperCount && trimRemark[len(trimRemark)-1] == '?' {
		return "Calm down, I know what I'm doing!"
	}
	if alphabetCount == upperCount {
		return "Whoa, chill out!"
	}
	if trimRemark[len(trimRemark)-1] == '?' {
		return "Sure."
	}
	return "Whatever."
}
