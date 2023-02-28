package protein

import "errors"

var trans = map[string]string {
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

var ErrStop = errors.New("ErrStop")
var ErrInvalidBase = errors.New("ErrInvalidBase")

func FromRNA(rna string) ([]string, error) {
	var out []string
	for i := 0; i < len(rna)/3; i++ {
		codon := rna[i*3 : i*3 + 3]
		prot, err := FromCodon(codon)
		if err == nil {
			out = append(out, prot)
		} else if errors.Is(err, ErrStop) {
			return out, nil
		} else {
			return nil, err
		}
	}
	return out, nil
}

func FromCodon(codon string) (string, error) {
	prot, ok := trans[codon]

	if !ok {
		return "", ErrInvalidBase
	}

	if prot == "STOP" {
		return "", ErrStop
	}

	return prot, nil
}
