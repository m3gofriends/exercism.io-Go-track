// Package wordy parses and evaluates simple math word problems returning the answer as an integer.
package wordy

import (
	"math"
	"strconv"
	"strings"
)

const (
	addition       string = "plus"
	subtraction    string = "minus"
	multiplication string = "multiplied"
	division       string = "divided"
	power          string = "raised"
	cube           string = "cubed"
)

// Answer parses and evaluates simple math word problems returning the answer as an integer.
func Answer(s string) (int, bool) {
	numberList, operationList, flag := []int{}, []string{}, false
	s = strings.ReplaceAll(s[:len(s)-1], "th", "")

	for _, v := range strings.Split(s, " ") {
		if number, err := strconv.Atoi(v); err == nil {
			if flag {
				return 0, false
			}
			numberList, flag = append(numberList, number), true
		} else if v == addition || v == subtraction || v == multiplication || v == division || v == power || v == cube {
			if !flag || v == cube {
				return 0, false
			}
			operationList, flag = append(operationList, v), false
		}
	}

	if len(numberList)-1 != len(operationList) {
		return 0, false
	}

	for i := 0; i < len(operationList); i++ {
		switch operationList[i] {
		case addition:
			numberList[0] += numberList[i+1]
		case subtraction:
			numberList[0] -= numberList[i+1]
		case multiplication:
			numberList[0] *= numberList[i+1]
		case division:
			numberList[0] /= numberList[i+1]
		default:
			numberList[0] = int(math.Pow(float64(numberList[0]), float64(numberList[i+1])))
		}
	}

	return numberList[0], true
}
