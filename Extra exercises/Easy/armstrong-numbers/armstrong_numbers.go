// Package armstrong checks if it is an Armstrong number.
package armstrong

import (
	"math"
	"strconv"
)

// IsNumber checks if it is an Armstrong number.
func IsNumber(input int) bool {
	inputByte, output := []byte(strconv.Itoa(input)), 0
	length := len(inputByte)

	for i := 0; i < length; i++ {
		output += int(math.Pow(float64(inputByte[i]-'0'), float64(length)))
	}

	return input == output
}
