package benchmarkfmt

import (
	"bytes"
	"fmt"
	"testing"
)

// A simple custom struct type
type Person struct {
	Name string
	Age  int
}

// A recursive custom struct type
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (p Person) String() string {
	return fmt.Sprintf("Person{Name:%s, Age:%d}", p.Name, p.Age)
}

func (n Node) String() string {
	if n.Left == nil && n.Right == nil {
		return fmt.Sprintf("%d", n.Value)
	}
	return fmt.Sprintf("(%s %s)", fmt.Sprint(n.Left), fmt.Sprint(n.Right))
}

// Benchmark for fmt.Sprintf with different data types
func BenchmarkSprintf(b *testing.B) {
	var (
		intVal    = 42
		floatVal  = 3.14159
		stringVal = "Hello, World!"
		personVal = Person{Name: "Alice", Age: 25}
		nodeVal   = Node{Value: 1, Left: &Node{Value: 2}, Right: &Node{Value: 3}}
	)

	// Benchmark with int
	b.Run("Int", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%d", intVal)
		}
	})

	// Benchmark with float
	b.Run("Float", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%.2f", floatVal)
		}
	})

	// Benchmark with string
	b.Run("String", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s", stringVal)
		}
	})

	// Benchmark with interface (Person)
	b.Run("Interface-Person", func(b *testing.B) {
		var value interface{} = personVal
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", value)
		}
	})

	// Benchmark with interface (Node)
	b.Run("Interface-Node", func(b *testing.B) {
		var value interface{} = nodeVal
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%v", value)
		}
	})

	// Benchmark with custom struct (Person)
	b.Run("CustomStruct-Person", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s", personVal)
		}
	})

	// Benchmark with custom struct (Node)
	b.Run("CustomStruct-Node", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = fmt.Sprintf("%s", nodeVal)
		}
	})
}

// Run the tests using "go test -bench=. -benchmem"
