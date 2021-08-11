// Package acronym converts phrase to acronym.
package acronym

import (
	"regexp"
	"strings"
)

var regex = regexp.MustCompile(`[A-Z']+`)

// Abbreviate converts phrase to acronym.
func Abbreviate(s string) (output string) {
	for _, v := range regex.FindAllString(strings.ToUpper(s), -1) {
		output += string(v[0])
	}
	return output
}
