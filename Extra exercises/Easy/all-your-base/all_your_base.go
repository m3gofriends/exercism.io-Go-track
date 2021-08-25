// Package allyourbase converts a number from one base to another base.
package allyourbase

import "errors"

// ConvertToBase converts a number from one base to another base.
func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (outputDigits []int, err error) {
	if inputBase < 2 {
		return []int(nil), errors.New("input base must be >= 2")
	} else if outputBase < 2 {
		return []int(nil), errors.New("output base must be >= 2")
	}

	sum := 0
	for i, j := len(inputDigits)-1, 1; i >= 0; i, j = i-1, j*inputBase {
		if inputDigits[i] < 0 || inputDigits[i] >= inputBase {
			return []int(nil), errors.New("all digits must satisfy 0 <= d < input base")
		}
		sum += inputDigits[i] * j
	}

	if sum == 0 {
		return []int{0}, err
	}

	for ; sum > 0; sum /= outputBase {
		outputDigits = append(outputDigits, (sum % outputBase))
	}

	return reverse(outputDigits), err
}

func reverse(r []int) []int {
	for i, j := 0, len(r)-1; i <= j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}
