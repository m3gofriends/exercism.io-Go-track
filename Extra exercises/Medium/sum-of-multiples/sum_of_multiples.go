// Package summultiples finds the sum of all unique multiples of a particular number.
package summultiples

type divisorStruct struct {
	divisor     int
	accumulator int
}

func (d *divisorStruct) accumulate() {
	d.accumulator += d.divisor
}

// SumMultiples finds the sum of all unique multiples of a particular number.
func SumMultiples(limit int, divisors ...int) (sum int) {
	d := []divisorStruct{}

	for _, divisor := range divisors {
		if divisor > 0 {
			d = append(d, divisorStruct{divisor, divisor})
		}
	}

	if len(d) == 0 {
		return sum
	}

	for {
		minValue, minIndex := d[0].accumulator, 0
		for i := 1; i < len(d); i++ {
			if minValue > d[i].accumulator {
				minValue = d[i].accumulator
				minIndex = i
			} else if minValue == d[i].accumulator {
				d[i].accumulate()
			}
		}
		if minValue >= limit {
			return sum
		}
		sum += minValue
		d[minIndex].accumulate()
	}
}
