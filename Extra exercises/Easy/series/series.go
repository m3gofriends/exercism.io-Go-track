// Package series implements a simple string handling.
package series

// All outputs all the contiguous substrings of length n in that string in the order that they appear.
func All(n int, s string) (output []string) {
	if n > len(s) {
		return nil
	}

	for i := 0; i+n-1 < len(s); i++ {
		output = append(output, s[i:i+n])
	}

	return output
}

// UnsafeFirst outputs the first substring of length n of the string.
func UnsafeFirst(n int, s string) string {
	return s[:n]
}
