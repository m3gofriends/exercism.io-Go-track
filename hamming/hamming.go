// Package hamming calculates hamming distance.
package hamming

import "errors"

// Distance calculates hamming distance.
func Distance(a, b string) (distance int, err error) {
	if len(a) != len(b) {
		return -1, errors.New("sequence's length are different")
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, err
}
