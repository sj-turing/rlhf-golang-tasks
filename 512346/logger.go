package main

import (
	"log"
	"math/rand"
	"os"
)

func main() {
	file, err := os.OpenFile("logs.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Fatalf("error while reading logs.txt file: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	i := 10
	for i > 0 {
		log.Println("printing random number", rand.Int63())
		i--
	}
}
