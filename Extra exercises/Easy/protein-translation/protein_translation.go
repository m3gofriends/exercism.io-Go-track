// Package protein implements a simple string conversion.
package protein

import "errors"

// Error codes
var (
	ErrStop        = errors.New("stop")
	ErrInvalidBase = errors.New("invalid base")
)

// FromCodon translates the codon sequence into a polypeptide.
func FromCodon(codon string) (polypeptide string, err error) {
	switch codon {
	case "AUG":
		return "Methionine", err
	case "UUU", "UUC":
		return "Phenylalanine", err
	case "UUA", "UUG":
		return "Leucine", err
	case "UCU", "UCC", "UCA", "UCG":
		return "Serine", err
	case "UAU", "UAC":
		return "Tyrosine", err
	case "UGU", "UGC":
		return "Cysteine", err
	case "UGG":
		return "Tryptophan", err
	case "UAA", "UAG", "UGA":
		return "", ErrStop
	default:
		return "", ErrInvalidBase
	}
}

// FromRNA translates RNA sequences into proteins.
func FromRNA(rna string) (protein []string, err error) {
	for i := 0; i < len(rna); i += 3 {
		codon, err := FromCodon(rna[i : i+3])
		switch err {
		case ErrStop:
			return protein, nil
		case ErrInvalidBase:
			return protein, ErrInvalidBase
		default:
			protein = append(protein, codon)
		}
	}
	return protein, err
}
