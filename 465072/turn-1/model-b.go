package main

import (
	"fmt"
)

// Point represents a point in 2D space
type Point struct {
	X int
	Y int
}

// String returns a custom formatted string representation of a Point
func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)", p.X, p.Y)
}

func main() {
	p := Point{X: 1, Y: 2}
	fmt.Println(p) // Output: (1, 2)

	// You can also use fmt.Sprintf to get the formatted string
	formattedString := fmt.Sprintf("The point is: %v", p)
	fmt.Println(formattedString) // Output: The point is: (1, 2)
}
