// Package variablelengthquantity implements variable length quantity encoding and decoding.
package variablelengthquantity

import (
	"errors"
	"fmt"
	"strconv"
)

// EncodeVarint implements variable length quantity encoding.
func EncodeVarint(list []uint32) (output []byte) {
	for i := len(list) - 1; i >= 0; i-- {
		s, index := decToBin(list[i]), 0
		for ; len(s) > 7; index++ {
			output, s = append(output, binToDec(s[len(s)-7:], index)), s[:len(s)-7]
		}

		switch value, _ := strconv.ParseInt(s, 2, 32); {
		case index == 0:
			output = append(output, byte(value))
		default:
			output = append(output, byte(128+value))
		}
	}

	return reverse(output)
}

// DecodeVarint implements variable length quantity decoding.
func DecodeVarint(list []byte) (output []uint32, err error) {
	s, outputString := "", ""
	for i := 0; i < len(list); i++ {
		if s = decToBin(uint32(list[i])); len(s) != 8 {
			for i := len(s); i < 7; i++ {
				s = "0" + s
			}
			outputString += s
			value, _ := strconv.ParseInt(outputString, 2, 32)
			output, outputString = append(output, uint32(value)), ""
		} else {
			outputString += s[1:]
		}
	}

	if len(s) == 8 {
		return []uint32{}, errors.New("incomplete sequence causes error")
	}

	return output, err
}

func decToBin(n uint32) (output string) {
	if n == 0 {
		return "0"
	}

	for ; n != 0; n /= 2 {
		output = fmt.Sprint(n%2) + output
	}

	return output
}

func binToDec(s string, index int) byte {
	value, _ := strconv.ParseInt(s, 2, 32)

	if index == 0 {
		return byte(value)
	}

	return byte(128 + value)
}

func reverse(r []byte) []byte {
	for i, j := 0, len(r)-1; i <= j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return r
}
