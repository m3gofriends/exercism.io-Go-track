// Package scale generates the musical scale starting with the tonic and following the specified interval pattern.
package scale

var (
	sharps = [12]string{"A", "A#", "B", "C", "C#", "D", "D#", "E", "F", "F#", "G", "G#"}
	flats  = [12]string{"A", "Bb", "B", "C", "Db", "D", "Eb", "E", "F", "Gb", "G", "Ab"}
)

// Scale generates the musical scale starting with the tonic and following the specified interval pattern.
func Scale(tonic string, interval string) (output []string) {
	switch tonic {
	case "F", "Bb", "Eb", "Ab", "Db", "Gb", "d", "g", "c", "f", "bb", "eb":
		tonicIndex := searchIndex(capitalize(tonic), flats)
		if interval == "" {
			for i := 0; i < 12; i++ {
				output = append(output, flats[tonicIndex])
				tonicIndex = (tonicIndex + 1) % 12
			}
		} else {
			for i := 0; i < len(interval); i++ {
				output = append(output, flats[tonicIndex])
				switch interval[i] {
				case 'm':
					tonicIndex = (tonicIndex + 1) % 12
				case 'M':
					tonicIndex = (tonicIndex + 2) % 12
				default:
					tonicIndex = (tonicIndex + 3) % 12
				}
			}
		}
	default:
		tonicIndex := searchIndex(capitalize(tonic), sharps)
		if interval == "" {
			for i := 0; i < 12; i++ {
				output = append(output, sharps[tonicIndex])
				tonicIndex = (tonicIndex + 1) % 12
			}
		} else {
			for i := 0; i < len(interval); i++ {
				output = append(output, sharps[tonicIndex])
				switch interval[i] {
				case 'm':
					tonicIndex = (tonicIndex + 1) % 12
				case 'M':
					tonicIndex = (tonicIndex + 2) % 12
				default:
					tonicIndex = (tonicIndex + 3) % 12
				}
			}
		}
	}
	return output
}

func capitalize(s string) string {
	if s[0] > 96 && s[0] < 123 { // a to z
		sByte := []byte(s)
		sByte[0] -= 32
		s = string(sByte)
	}
	return s
}

func searchIndex(s string, slice [12]string) (index int) {
	for {
		if s == slice[index] {
			return index
		}
		index++
	}
}
