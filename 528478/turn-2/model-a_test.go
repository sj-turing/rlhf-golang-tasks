package benchmarkfmt

import (
	"bytes"
	"fmt"
	"testing"
)

// Custom struct type
type Person struct {
	Name string
	Age  int
}

// Recursive type (basic linked list node)
type Node struct {
	Value int
	Next  *Node
}

// Benchmark for fmt.Sprintf with various data types
func BenchmarkSprintf(b *testing.B) {
	var (
		intVal                   = 42
		floatVal                 = 3.14159
		stringVal                = "Hello, World!"
		interfaceVal interface{} = "Interface String"
		person                   = Person{Name: "Alice", Age: 30}
		node                     = &Node{Value: 1, Next: &Node{Value: 2, Next: nil}}
	)

	b.Run("Int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%d", intVal)
		}
	})

	b.Run("Float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%.2f", floatVal)
		}
	})

	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s", stringVal)
		}
	})

	b.Run("Interface", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", interfaceVal)
		}
	})

	b.Run("Struct", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%+v", person)
		}
	})

	b.Run("Recursive", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%+v", node)
		}
	})
}

// Benchmark for fmt.Printf with various data types using a buffer
func BenchmarkPrintf(b *testing.B) {
	var (
		intVal                   = 42
		floatVal                 = 3.14159
		stringVal                = "Hello, World!"
		interfaceVal interface{} = "Interface String"
		person                   = Person{Name: "Alice", Age: 30}
		node                     = &Node{Value: 1, Next: &Node{Value: 2, Next: nil}}
	)

	b.Run("Int", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset() // Reset buffer to avoid accumulation
			fmt.Fprintf(&buf, "%d", intVal)
		}
	})

	b.Run("Float", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			fmt.Fprintf(&buf, "%.2f", floatVal)
		}
	})

	b.Run("String", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			fmt.Fprintf(&buf, "%s", stringVal)
		}
	})

	b.Run("Interface", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			fmt.Fprintf(&buf, "%v", interfaceVal)
		}
	})

	b.Run("Struct", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			fmt.Fprintf(&buf, "%+v", person)
		}
	})

	b.Run("Recursive", func(b *testing.B) {
		var buf bytes.Buffer
		for i := 0; i < b.N; i++ {
			buf.Reset()
			fmt.Fprintf(&buf, "%+v", node)
		}
	})
}
