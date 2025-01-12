package main

import (
	"fmt"
	"strings"
)

func main() {
	var sb strings.Builder
	for i := 0; i < 1000000; i++ {
		sb.WriteString("Hello, World! ")
	}
	fmt.Println(sb.String())
}
