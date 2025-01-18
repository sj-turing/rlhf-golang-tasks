package main

import (
	"bytes"
	"math/rand"
	"testing"
)

func DNAMatching(dna1, dna2 string) bool {
	// this is not actual algorithm
	// just mimicing the DNA matching algorithm
	// to show the usage of pprof

	if len(dna1) != len(dna2) {
		return false
	}

	for index := range dna1 {
		if dna1[index] != dna2[index] {
			return false
		}
	}
	return true
}

const dnaCode string = "ACGT"

func getRandomDNASequence(size int) string {
	var seq = &bytes.Buffer{}
	defer seq.Reset()
	for size > 0 {
		seq.WriteByte(dnaCode[rand.Intn(len(dnaCode))])
		size--
	}

	return seq.String()
}

func BenchmarkDNAMatching(b *testing.B) {
	dataset := []struct {
		dna1, dna2 string
	}{
		{
			dna1: getRandomDNASequence(rand.Intn(10000)),
			dna2: getRandomDNASequence(rand.Intn(10000)),
		},
		{
			dna1: getRandomDNASequence(rand.Intn(20000)),
			dna2: getRandomDNASequence(rand.Intn(20000)),
		},
		{
			dna1: getRandomDNASequence(rand.Intn(25000)),
			dna2: getRandomDNASequence(rand.Intn(25000)),
		},
		{
			dna1: getRandomDNASequence(rand.Intn(30000)),
			dna2: getRandomDNASequence(rand.Intn(30000)),
		},
		{
			dna1: getRandomDNASequence(rand.Intn(35000)),
			dna2: getRandomDNASequence(rand.Intn(35000)),
		},
	}
	for i := 0; i < b.N; i++ {
		mod := (i + 1) % len(dataset)
		data := dataset[mod]
		DNAMatching(data.dna1, data.dna2)
	}
}
