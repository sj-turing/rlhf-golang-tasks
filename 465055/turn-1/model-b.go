package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	enc, _ := json.Marshal(p)
	output := fmt.Sprintf("Formatted JSON: %s\n", string(enc))
	fmt.Print(output)
}
