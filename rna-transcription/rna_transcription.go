// Package strand implements a simple string conversion.
package strand

// ToRNA converts Dna strings to Rna strings.
func ToRNA(dna string) string {
	byteDna := []byte(dna)
	for i := 0; i < len(byteDna); i++ {
		switch string(byteDna[i]) {
		case "G":
			byteDna[i] = 'C'
		case "C":
			byteDna[i] = 'G'
		case "T":
			byteDna[i] = 'A'
		case "A":
			byteDna[i] = 'U'
		}
	}
	return string(byteDna)
}
