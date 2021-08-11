// Package darts returns the earned points in a single toss of a Darts game.
package darts

import "math"

// Score returns the earned points in a single toss of a Darts game.
func Score(x, y float64) int {
	switch distance := math.Hypot(x, y); {
	case distance <= 1:
		return 10
	case distance <= 5:
		return 5
	case distance <= 10:
		return 1
	default:
		return 0
	}
}
