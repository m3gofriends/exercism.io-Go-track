// Package bob classifies the message.
package bob

import (
	"regexp"
	"strings"
)

// Hey classifies the message.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	upperCase, _ := regexp.MatchString(`^[^a-z]*[A-Z]+[^a-z]*$`, remark)
	questionMark, _ := regexp.MatchString(`\?$`, remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	if upperCase && questionMark {
		return "Calm down, I know what I'm doing!"
	}
	if upperCase {
		return "Whoa, chill out!"
	}
	if questionMark {
		return "Sure."
	}
	return "Whatever."
}
