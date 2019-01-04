// Package strand contains utilities for RNA strings
package strand

var rnaFlip = map[rune]string{
	'G': "C",
	'C': "G",
	'T': "A",
	'A': "U",
}

// ToRNA takes an string representing dna and transforms it into its RNA
// equiavalent
func ToRNA(dna string) string {
	var flippedDna string

	for _, char := range dna {
		flippedDna += rnaFlip[char]
	}

	return flippedDna
}
