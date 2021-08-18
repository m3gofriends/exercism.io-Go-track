// Package cryptosquare implements the classic method for composing secret messages.
package cryptosquare

import (
	"math"
	"strings"
)

// Encode implements the classic method for composing secret messages.
func Encode(s string) string {
	if len(s) < 2 {
		return s
	}

	var c, r, floor, ceil int

	s = strings.ToLower(s)
	for i := len(s) - 1; i > -1; i-- {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			s = s[:i] + s[i+1:]
		}
	}

	floor = int(math.Floor(math.Sqrt(float64(len(s)))))
	ceil = int(math.Ceil(math.Sqrt(float64(len(s)))))

	switch length := len(s); {
	case floor*floor >= length:
		c, r = floor, floor
	case (floor+1)*floor >= length:
		c, r = floor+1, floor
	case ceil*ceil >= length:
		c, r = ceil, ceil
	case (ceil+1)*ceil >= length:
		c, r = ceil, ceil
	}

	return string(normalize(s, c, r))
}

func normalize(s string, c, r int) (output []byte) {
	for i := 0; i < c; i++ {
		for index := i; index < c*r; index += c {
			if index < len(s) {
				output = append(output, s[index])
			} else {
				output = append(output, ' ')
			}
		}
		output = append(output, ' ')
	}
	return output[:len(output)-1]
}
