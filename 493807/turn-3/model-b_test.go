package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkFormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(987654321, 10)
	}
}

func BenchmarkFormatFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatFloat(3.14, 'f', 2, 64)
	}
}

func BenchmarkItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(12345)
	}
}

func BenchmarkReplace(b *testing.B) {
	input := "Hello, World! Hello, Go!"
	for i := 0; i < b.N; i++ {
		strings.Replace(input, "Hello, ", "Hi, ", -1)
	}
}

func BenchmarkFscanf(b *testing.B) {
	input := bytes.NewBufferString("Alice 95\nBob 88\nCharlie 76")
	for i := 0; i < b.N; i++ {
		var name string
		var score int
		for {
			_, err := fmt.Fscanf(input, "%s %d\n", &name, &score)
			if err != nil {
				break
			}
		}
		input.Reset()
	}
}
