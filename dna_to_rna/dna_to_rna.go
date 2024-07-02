package dna_to_rna

func DNAtoRNA(dna string) string {
	var rna string = ""

	for _, char := range dna {
		if char == 'T' {
			rna += "U"
		} else {
			rna += string(char)
		}
	}

	return rna
}
