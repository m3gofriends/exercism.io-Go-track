// Package triangle calculates the type of triangle.
package triangle

import "math"

// Kind is an integer.
type Kind int

// The integer returned according to the type of the triangle.
const (
	NaT = iota
	Equ
	Iso
	Sca
)

// KindFromSides return the type of the triangle.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if (a <= 0 || b <= 0 || c <= 0) || math.IsInf(a+b+c, 0) {
		k = NaT
	} else if a+b >= c && a+c >= b && b+c >= a {
		if a == b && b == c {
			k = Equ
		} else if a == b || a == c || b == c {
			k = Iso
		} else {
			k = Sca
		}
	}
	return k
}
