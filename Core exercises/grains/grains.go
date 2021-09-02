// Package grains implements a simple mathematical calculation.
package grains

import "errors"

// Square calculates the number of grains of wheat on the chessboard.
func Square(n int) (uint64, error) {
	if n < 1 || n > 64 {
		return 0, errors.New("out of range")
	}

	return 1 << (n - 1), nil
}

// Total calculates the number of all the grains of wheat on the chessboard.
func Total() uint64 {
	return 1<<64 - 1
}
