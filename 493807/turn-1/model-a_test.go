package main

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Hello, %v! Your score is %d.", "Alice", 95)
	}
}

func BenchmarkConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "Hello, " + "Alice" + "! Your score is " + fmt.Sprint(95) + "."
	}
}

func BenchmarkTemplate(b *testing.B) {
	t := template.Must(template.New("greeting").Parse(`Hello, {{ .Name }}! Your score is {{ .Score }}.`))
	data := struct {
		Name  string
		Score int
	}{
		Name:  "Alice",
		Score: 95,
	}

	for i := 0; i < b.N; i++ {
		var message bytes.Buffer
		_ = t.Execute(&message, data)
	}
}
