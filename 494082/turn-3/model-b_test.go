package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func joinWordsFmtSprintf(words []string, delimiter string) string {
	return fmt.Sprintf("%s", words)
}

func BenchmarkJoinWordsFmtSprintf(b *testing.B) {
	words := []string{"Hello", "World"}
	for n := 0; n < b.N; n++ {
		joinWordsFmtSprintf(words, " ")
	}
}

func joinWordsStringsJoin(words []string, delimiter string) string {
	return strings.Join(words, delimiter)
}

func BenchmarkJoinWordsStringsJoin(b *testing.B) {
	words := []string{"Hello", "World"}
	for n := 0; n < b.N; n++ {
		joinWordsStringsJoin(words, " ")
	}
}

func joinWordsBytesBuffer(words []string, delimiter string) string {
	var buffer bytes.Buffer
	for i, word := range words {
		if i > 0 {
			buffer.WriteString(delimiter)
		}
		buffer.WriteString(word)
	}
	return buffer.String()
}

func BenchmarkJoinWordsBytesBuffer(b *testing.B) {
	words := []string{"Hello", "World"}
	for n := 0; n < b.N; n++ {
		joinWordsBytesBuffer(words, " ")
	}
}

func joinWordsCustom(words []string, delimiter string) string {
	result := ""
	for i, word := range words {
		if i > 0 {
			result += delimiter
		}
		result += word
	}
	return result
}

func BenchmarkJoinWordsCustom(b *testing.B) {
	words := []string{"Hello", "World"}
	for n := 0; n < b.N; n++ {
		joinWordsCustom(words, " ")
	}
}
