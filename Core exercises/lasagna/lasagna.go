// Package lasagna implements a simple mathematical calculation.
package lasagna

// OvenTime defines the expected oven time in minutes.
func OvenTime() int {
	return 40
}

// RemainingOvenTime calculates the remaining oven time in minutes.
func RemainingOvenTime(n int) int {
	return 40 - n
}

// PreparationTime calculates the preparation time in minutes.
func PreparationTime(layers int) int {
	return layers * 2
}

// ElapsedTime calculates the elapsed working time in minutes.
func ElapsedTime(layers, total int) int {
	return PreparationTime(layers) + total
}
