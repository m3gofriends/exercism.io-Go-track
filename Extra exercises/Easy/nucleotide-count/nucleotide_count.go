// Package dna implements a simple DNA strand frequency count.
package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
type Histogram map[byte]int

// DNA is a string of nucleotides.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
func (d DNA) Counts() (histogram Histogram, err error) {
	histogram = make(map[byte]int)
	histogram['A'] = strings.Count(string(d), "A")
	histogram['C'] = strings.Count(string(d), "C")
	histogram['G'] = strings.Count(string(d), "G")
	histogram['T'] = strings.Count(string(d), "T")
	if histogram['A'] + histogram['C'] + histogram['G'] + histogram['T'] == len(d) {
		return histogram, err
	}
	return histogram, errors.New("strand with invalid nucleotides")
}
