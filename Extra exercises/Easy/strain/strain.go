// Package strain implements the keep and discard operation on collections.
package strain

type (
	// Ints is a collection of integers.
	Ints []int

	// Lists is a collection of Ints.
	Lists [][]int

	// Strings is a collection of strings.
	Strings []string
)

// Keep filters integers in a collection using the given predicate.
func (i Ints) Keep(pred func(int) bool) (output Ints) {
	for _, value := range i {
		if pred(value) {
			output = append(output, value)
		}
	}
	return output
}

// Discard filters out integers in a collection using the given predicate.
func (i Ints) Discard(pred func(int) bool) (output Ints) {
	for _, value := range i {
		if !pred(value) {
			output = append(output, value)
		}
	}
	return output
}

// Keep filters lists in a collection using the given predicate.
func (l Lists) Keep(pred func([]int) bool) (output Lists) {
	for _, value := range l {
		if pred(value) {
			output = append(output, value)
		}
	}
	return output
}

// Keep filters strings in a collection using the given predicate.
func (s Strings) Keep(pred func(string) bool) (output Strings) {
	for _, value := range s {
		if pred(value) {
			output = append(output, value)
		}
	}
	return output
}
