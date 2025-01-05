package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for _, value := range slice {
		fmt.Println(value)
	}
}
