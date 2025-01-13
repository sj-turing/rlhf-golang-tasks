package main

import (
	"testing"
)

func TestSequencialJoin(t *testing.T) {
	type args struct {
		strList []string
	}

	testCases := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "empty slice should returns empty string",
			args: args{
				strList: []string{},
			},
			expected: "",
		},
		{
			name: "slice of strings should get join with space as delimiter",
			args: args{
				strList: []string{"abc", "xyz", "pqr"},
			},
			expected: "abc xyz pqr",
		},
		{
			name: "slice with one element should not append any space to string",
			args: args{
				strList: []string{"abc"},
			},
			expected: "abc",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := sequentialJoin(tt.args.strList)
			if tt.expected != got {
				t.Fatalf("Expected %s, got %s", tt.expected, got)
			}
		})
	}
}

func TestConcurrentJoin(t *testing.T) {
	type args struct {
		strList []string
	}

	testCases := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "empty slice should returns empty string",
			args: args{
				strList: []string{},
			},
			expected: "",
		},
		{
			name: "slice of strings should get join with space as delimiter",
			args: args{
				strList: []string{"abc", "xyz", "pqr"},
			},
			expected: "abc xyz pqr",
		},
		{
			name: "slice with one element should not append any space to string",
			args: args{
				strList: []string{"abc"},
			},
			expected: "abc",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			got := concurrentJoin(tt.args.strList, 3)
			if tt.expected != got {
				t.Fatalf("Expected %s, got %s", tt.expected, got)
			}
		})
	}
}

func TestConcurrentJoinWithChannels(t *testing.T) {
	// as concurrent process does not garauntee sequence
	// for test case we can again split the result and check every string
	// existed there in the split slice
}

func BenchmarkSequentialJoin(b *testing.B) {
	strList := generateStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sequentialJoin(strList)
	}
}

func BenchmarkConcurrentJoin(b *testing.B) {
	strList := generateStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrentJoin(strList, len(strList))
	}
}

func BenchmarkConcurrentJoinWithChannel(b *testing.B) {
	strList := generateStrings(1000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		concurrentJoinWithChannels(strList)
	}
}
