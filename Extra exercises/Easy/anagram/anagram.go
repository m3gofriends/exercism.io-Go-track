// Package anagram checks the rearrangement of letters to form a new word.
package anagram

import "strings"

// Detect checks the rearrangement of letters to form a new word.
func Detect(subject string, candidates []string) (output []string) {
	subject = strings.ToLower(subject)

	for _, candidate := range candidates {
		if compare(subject, candidate) {
			output = append(output, candidate)
		}

	}
	return output
}

func compare(s, c string) bool {
	c = strings.ToLower(c)
	if s == c || len(s) != len(c) {
		return false
	}
	for _, letter := range c {
		i := strings.IndexRune(s, letter)
		if i == -1 {
			return false
		}
		s = s[:i] + s[i+1:]
	}
	return true
}
