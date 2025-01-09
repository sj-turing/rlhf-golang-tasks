package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type User struct {
	ID   int
	Name string
}

func main() {
	var users []User
	numUsers := 1000000
	numInsertions := 50000
	numDeletions := 50000
	fmt.Println("Starting simulation...")
	start := time.Now()
	// Insertions
	for i := 0; i < numInsertions; i++ {
		user := User{ID: i, Name: fmt.Sprintf("User %d", i)}
		users = append(users, user)
	}
	// Deletions
	for i := 0; i < numDeletions; i++ {
		index := rand.Intn(len(users))
		users = append(users[:index], users[index+1:]...)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Simulation completed in: %s\n", elapsed)
	// Force garbage collection to analyze the impact
	runtime.GC()
	fmt.Println("Garbage collection forced.")
	start = time.Now()
	// Insertions again to measure the impact of GC
	for i := numInsertions; i < numInsertions*2; i++ {
		user := User{ID: i, Name: fmt.Sprintf("User %d", i)}
		users = append(users, user)
	}
	end = time.Now()
	elapsed = end.Sub(start)
	fmt.Printf("Insertions after GC completed in: %s\n", elapsed)
}
