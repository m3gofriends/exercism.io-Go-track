// Package romannumerals converts number to Roman numeral.
package romannumerals

import "errors"

var (
	ones      = []string{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens      = []string{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}
	hundreds  = []string{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}
	thousands = []string{"", "M", "MM", "MMM"}
)

// ToRomanNumeral converts number to Roman numeral.
func ToRomanNumeral(number int) (romanNumeral string, err error) {
	if number < 1 || number > 3000 {
		return romanNumeral, errors.New("out of range")
	}
	romanNumeral = thousands[number/1000] + hundreds[number%1000/100] + tens[number%100/10] + ones[number%10]
	return romanNumeral, err
}
