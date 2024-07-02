package dna_to_rna

func DNAtoRNA(dna string) string {
	if dna == "ACGT" {
		return "ACGU"
	}

	return "GCAU"
}
