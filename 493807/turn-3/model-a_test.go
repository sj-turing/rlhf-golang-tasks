package main

/*

import (
	"bytes"
	"fmt"
	"html/template"
	"strings"
	"testing"
)

// Simple message with name and score
var message = "Hello, %v! Your score is %d."
var (
	name  = "Alice"
	score = 95
	parts = []string{"Hello, ", name, "! Your score is ", fmt.Sprint(score), "."}
	t     = template.Must(template.New("greeting").Parse(`Hello, {{ .Name }}! Your score is {{ .Score }}.`))
)

func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf(message, name, score)
	}
}

func BenchmarkConcatenation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = "Hello, " + name + "! Your score is " + fmt.Sprint(score) + "."
	}
}

func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Join(parts, "")
	}
}

func BenchmarkBytesBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		buf.WriteString("Hello, ")
		buf.WriteString(name)
		buf.WriteString("! Your score is ")
		buf.WriteString(fmt.Sprint(score))
		buf.WriteString(".")
		_ = buf.String()
	}
}

func BenchmarkTemplate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		data := struct {
			Name  string
			Score int
		}{
			Name:  name,
			Score: score,
		}
		if err := t.Execute(&buf, data); err != nil {
			b.Fatalf("template execute failed: %v", err)
		}
	}
}
*/
