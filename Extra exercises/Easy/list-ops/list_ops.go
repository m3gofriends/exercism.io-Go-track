// Package listops implements a series of basic list operations, without using existing functions.
package listops

type (
	// IntList is an integer list
	IntList   []int
	binFunc   func(int, int) int
	predFunc  func(int) bool
	unaryFunc func(int) int
)

// Foldl folds each item into the accumulator from the left.
func (i IntList) Foldl(fn binFunc, initial int) (sum int) {
	if len(i) == 0 {
		return initial
	}
	for index := 0; index < len(i); index += 2 {
		sum += fn(i[index], i[index+1])
	}
	return fn(sum, initial)
}

// Foldr folds each item into the accumulator from the right.
func (i IntList) Foldr(fn binFunc, initial int) (sum int) {
	if len(i) == 0 {
		return initial
	}
	for index := 0; index < len(i); index += 2 {
		sum += fn(i[index+1], i[index])
	}
	return fn(initial, sum)
}

// Filter returns the list of all items for which predicate is true.
func (i IntList) Filter(fn predFunc) IntList {
	output := make(IntList, 0)
	for _, value := range i {
		if fn(value) {
			output = append(output, value)
		}
	}
	return output
}

// Length returns the number of elements in the list.
func (i IntList) Length() int {
	return len(i)
}

// Map applies the function to each item in the list.
func (i IntList) Map(fn unaryFunc) IntList {
	output := make(IntList, 0)
	for _, value := range i {
		output = append(output, fn(value))
	}
	return output
}

// Reverse reverses the elements in the list in place.
func (i IntList) Reverse() IntList {
	output := make(IntList, 0)
	for index, length := 0, len(i)-1; index < len(i); index++ {
		output = append(output, i[length-index])
	}
	return output
}

// Append extends the list with one or more elements.
func (i IntList) Append(input IntList) IntList {
	return append(i, input...)
}

// Concat combines one or more lists into one.
func (i IntList) Concat(input []IntList) IntList {
	for _, list := range input {
		i = append(i, list...)
	}
	return i
}
