package dnamatch

import (
	"testing"
)

func TestMatchPercentage(t *testing.T) {
	tests := []struct {
		seq1    string
		seq2    string
		percent float64
	}{
		{seq1: "ATCG", seq2: "ATCG", percent: 100.0},
		{seq1: "ATCG", seq2: "TAGC", percent: 50.0},
		{seq1: "ATCG", seq2: "AAGC", percent: 75.0},
		{seq1: "ATCG", seq2: "CCCC", percent: 0.0},
	}

	for _, test := range tests {
		percent := MatchPercentage(test.seq1, test.seq2)
		if percent != test.percent {
			t.Errorf("MatchPercentage(%q, %q) = %f, want %f", test.seq1, test.seq2, percent, test.percent)
		}
	}
}
func BenchmarkMatchPercentage(b *testing.B) {
	seq1 := "ATCG"
	seq2 := "TAGC"
	for i := 0; i < b.N; i++ {
		MatchPercentage(seq1, seq2)
	}
}
