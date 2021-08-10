// Package etl transforms the score system.
package etl

import "strings"

// Transform transforms the score system.
func Transform(oldScore map[int][]string) map[string]int {
	newScore := make(map[string]int)
	for key, value := range oldScore {
		for index := 0; index < len(value); index++ {
			newScore[strings.ToLower(value[index])] = key
		}
	}
	return newScore
}
