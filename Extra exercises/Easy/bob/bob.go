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

	switch {
	case len(remark) == 0:
		return "Fine. Be that way!"
	case upperCase && questionMark:
		return "Calm down, I know what I'm doing!"
	case upperCase:
		return "Whoa, chill out!"
	case questionMark:
		return "Sure."
	default:
		return "Whatever."
	}
}
