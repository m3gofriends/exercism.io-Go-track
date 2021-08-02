// Package twofer implements a simple message output.
package twofer

// ShareWith converts name to message.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
