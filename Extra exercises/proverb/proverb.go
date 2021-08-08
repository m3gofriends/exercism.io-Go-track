// Package proverb implements a simple message output.
package proverb

// Proverb converts string into slice.
func Proverb(rhyme []string) (output []string) {
	if len(rhyme) > 0 {
		for i := 0; i < len(rhyme)-1; i++ {
			output = append(output, "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost.")
		}
		output = append(output, "And all for the want of a " + rhyme[0] + ".")
	}
	return output
}
