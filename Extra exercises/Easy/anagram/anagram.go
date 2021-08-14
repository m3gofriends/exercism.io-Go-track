// Package anagram checks the rearrangement of indexs to form a new word.
package anagram

import "strings"

// Detect checks the rearrangement of indexs to form a new word.
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
	for i := 0; i < len(c); i++ {
		index := strings.IndexByte(s, c[i])
		if index == -1 {
			return false
		}
		s = s[:index] + s[index+1:]
	}
	return true
}
