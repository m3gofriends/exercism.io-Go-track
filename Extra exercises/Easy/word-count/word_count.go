// Package wordcount calculates the number of occurrences of each word in the phrase.
package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is a map.
type Frequency map[string]int

var regex = regexp.MustCompile(`[a-z0-9]+'?[a-z]?\b`)

// WordCount calculates the number of occurrences of each word in the phrase.
func WordCount(s string) Frequency {
	frequency := make(Frequency)

	s = strings.ToLower(strings.ReplaceAll(s, "\n", " "))
	for _, value := range regex.FindAllString(s, -1) {
		frequency[value]++
	}
	return frequency
}
