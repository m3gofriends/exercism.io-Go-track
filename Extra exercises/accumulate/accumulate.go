// Package accumulate implements the slice operation.
package accumulate

// Accumulate implements the slice operation.
func Accumulate(s []string, funcString func(string) string) (output []string) {
	for _, v := range s {
		output = append(output, funcString(v))
	}
	return output
}
