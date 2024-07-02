package dna_to_rna

import "testing"

type TestCase struct {
	Name string
	Input string
	Want string
}

func TestDNAtoRNA(t *testing.T) {
	testCases := []TestCase{
		{Name: "GCAT should return GCAU", Input: "GCAT", Want: "GCAU"},
		{Name: "ACGT should return ACGU", Input: "ACGT", Want: "ACGU"},
		{Name: "CGTA should return CGUA", Input: "CGTA", Want: "CGUA"},
	}

	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			got := DNAtoRNA(tt.Input)
			if got != tt.Want {
				t.Errorf("[ERROR] %v: got %v, but want %v", tt.Name, got, tt.Want)
			}
		})
	}
}
