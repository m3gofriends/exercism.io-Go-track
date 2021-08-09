// Package scrabble converts string to score.
package scrabble

import "strings"

// Score converts string to score.
func Score(s string) (score int) {
	for _, value := range s {
		switch strings.ToUpper(string(value)) {
		case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
			score++
		case "D", "G":
			score += 2
		case "B", "C", "M", "P":
			score += 3
		case "F", "H", "V", "W", "Y":
			score += 4
		case "K":
			score += 5
		case "J", "X":
			score += 8
		case "Q", "Z":
			score += 10
		}
	}
	return score
}
