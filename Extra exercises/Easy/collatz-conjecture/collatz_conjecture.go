// Package collatzconjecture implements a simple mathematical calculation.
package collatzconjecture

import "errors"

// CollatzConjecture implements a simple mathematical calculation.
func CollatzConjecture(n int) (counter int, err error) {
	if n < 1 {
		return counter, errors.New("out of range")
	}
	for ; n != 1; counter++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	return counter, err
}
