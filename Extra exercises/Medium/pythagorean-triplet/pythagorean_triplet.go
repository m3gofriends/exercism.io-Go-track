// Package pythagorean finds the Pythagorean triplets.
package pythagorean

import "sort"

// Triplet represents a Pythagorean triplet.
type Triplet [3]int

var pythagoreanTriple []Triplet = generatePythagorean(426)

// Range returns the list of Pythagorean triplets with sides in the range min to max inclusive.
func Range(min, max int) (output []Triplet) {
	for i := 0; i < len(pythagoreanTriple); i++ {
		if pythagoreanTriple[i][0] >= min && pythagoreanTriple[i][2] <= max {
			output = append(output, pythagoreanTriple[i])
		}
	}

	return output
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c is equal to p.
func Sum(p int) (output []Triplet) {
	for i := 0; i < len(pythagoreanTriple); i++ {
		if pythagoreanTriple[i][0]+pythagoreanTriple[i][1]+pythagoreanTriple[i][2] == p {
			output = append(output, pythagoreanTriple[i])
		}
	}

	return output
}

func generatePythagorean(n int) (pythagoreanList []Triplet) {
	pythagoreanMap := make(map[Triplet]bool)

	for i := 2; i*i+1 < n; i++ {
		for k := 1; k*(i*i+1) < n; k++ {
			value := Triplet{k * (i*i - 1), k * (2 * i), k * (i*i + 1)}.sort() // k*(i^2-j^2), k*(2*i*j), k*(i^2+j^2)
			if _, ok := pythagoreanMap[value]; !ok {
				pythagoreanMap[value] = true
				pythagoreanList = append(pythagoreanList, value)
			}
		}
	}

	sort.Slice(pythagoreanList, func(i, j int) bool { return pythagoreanList[i][0] < pythagoreanList[j][0] })
	return pythagoreanList
}

func (t Triplet) sort() Triplet {
	if t[0] > t[1] {
		t[0], t[1] = t[1], t[0]
	}
	return t
}
