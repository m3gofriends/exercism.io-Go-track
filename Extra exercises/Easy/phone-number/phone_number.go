// Package phonenumber provides phone number verification for NANP(North American Numbering Plan).
package phonenumber

import (
	"errors"
	"fmt"
)

// Number cleans up the input string and returns the 10-digit phone number.
func Number(s string) (string, error) {
	sByte := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] > 47 && s[i] < 58 { // 0 to 9
			sByte = append(sByte, s[i])
		}
	}

	if len(sByte) == 11 && sByte[0] == '1' {
		sByte = sByte[1:]
	}

	switch {
	case len(sByte) != 10:
		return "", errors.New("invalid length")
	case sByte[0] < '2':
		return "", errors.New("invalid area code")
	case sByte[3] < '2':
		return "", errors.New("invalid exchange code")
	}

	return string(sByte), nil
}

// AreaCode returns the area code from a NANP number.
func AreaCode(s string) (output string, err error) {
	output, err = Number(s)

	if err == nil {
		return output[:3], err
	}

	return "", err
}

// Format returns a consistently formatted NANP number.
func Format(s string) (output string, err error) {
	output, err = Number(s)

	if err == nil {
		return fmt.Sprintf("(%s) %s-%s", output[:3], output[3:6], output[6:]), err
	}

	return "", err
}
