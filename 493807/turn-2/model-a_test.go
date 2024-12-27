package main

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
)

func BenchmarkStringConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "Hello, " + "Alice" + "! Your score is " + fmt.Sprint(95) + "."
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("Hello, %v! Your score is %d.", "Alice", 95)
	}
}

func BenchmarkStringPercent(b *testing.B) {
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

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.WriteString("Hello, ")
		buf.WriteString("Alice")
		buf.WriteString("! Your score is ")
		fmt.Fprint(&buf, 95)
		buf.WriteString(".")
		_ = buf.String()
	}
}
