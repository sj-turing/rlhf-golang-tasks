package main

import (
	"fmt"
	"math"
	"runtime"
	"testing"
	"time"
)

func BenchmarkPrintfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("%d", i)
	}
}

func BenchmarkSprintfInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%d", i)
	}
}

func BenchmarkPrintfFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("%f", math.Pi)
	}
}

func BenchmarkSprintfFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%f", math.Pi)
	}
}

func BenchmarkPrintfString(b *testing.B) {
	s := "Hello, World!"
	for i := 0; i < b.N; i++ {
		fmt.Printf("%s", s)
	}
}

func BenchmarkSprintfString(b *testing.B) {
	s := "Hello, World!"
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%s", s)
	}
}

func BenchmarkPrintfStruct(b *testing.B) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 25}
	for i := 0; i < b.N; i++ {
		fmt.Printf("%+v", p)
	}
}

func BenchmarkSprintfStruct(b *testing.B) {
	type Person struct {
		Name string
		Age  int
	}
	p := Person{Name: "Alice", Age: 25}
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("%+v", p)
	}
}

func TestMain(m *testing.M) {
	runtime.GOMAXPROCS(runtime.NumCPU())
	m.Run()
}
