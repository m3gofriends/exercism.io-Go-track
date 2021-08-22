// Package secret returns a sequence of events for the secret handshake.
package secret

// Handshake returns a sequence of events for the secret handshake.
func Handshake(n uint) (output []string) {
	if n&16 == 0 {
		if n&1 != 0 {
			output = append(output, "wink")
		}
		if n&2 != 0 {
			output = append(output, "double blink")
		}
		if n&4 != 0 {
			output = append(output, "close your eyes")
		}
		if n&8 != 0 {
			output = append(output, "jump")
		}
	} else {
		if n&8 != 0 {
			output = append(output, "jump")
		}
		if n&4 != 0 {
			output = append(output, "close your eyes")
		}
		if n&2 != 0 {
			output = append(output, "double blink")
		}
		if n&1 != 0 {
			output = append(output, "wink")
		}
	}

	return output
}
