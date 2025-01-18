package main

import "testing"

func TestDNAMatchEqual(t *testing.T) {
	match := DNAMatch("AGCT", "AGCT")
	if !match {
		t.Error("Expected DNA sequences to match")
	}
}

func TestDNAMatchNotEqual(t *testing.T) {
	match := DNAMatch("AGCT", "AGGT")
	if match {
		t.Error("Expected DNA sequences not to match")
	}
}

func TestDNAMatchDifferentLengths(t *testing.T) {
	match := DNAMatch("AGCT", "AGCTAG")
	if match {
		t.Error("Expected DNA sequences not to match due to different lengths")
	}
}

func BenchmarkDNAMatch(b *testing.B) {
	dna1 := "AGCTAGCTAGCTAGCTAGCT"
	dna2 := "AGCTAGCTAGCTAGCTAGCT"

	for i := 0; i < b.N; i++ {
		DNAMatch(dna1, dna2)
	}
}
