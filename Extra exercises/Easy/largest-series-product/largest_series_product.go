// Package lsproduct calculates the largest product for a contiguous substring of digits of length n.
package lsproduct

import "errors"

// LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(s string, span int) (int64, error) {
	if len(s) < span {
		return -1, errors.New("span must be smaller than string length")
	} else if span < 0 {
		return -1, errors.New("span must be greater than zero")
	}

	var nowValue, product int64 = 1, 0
	for i, times := 0, len(s)-span+1; i < times; i, nowValue = i+1, 1 {
		for j := 0; j < span; j++ {
			if s[i+j] < '0' || s[i+j] > '9' {
				return -1, errors.New("digits input must only contain digits")
			}
			nowValue *= int64(s[i+j] - '0')
		}
		if nowValue > product {
			product = nowValue
		}
	}

	return product, nil
}
