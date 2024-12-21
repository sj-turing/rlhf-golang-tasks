package main

import (
	"fmt"
	"sync"
)

type Resource struct {
	mu sync.Mutex
	id int
}

var (
	resources   = make(map[int]*Resource)
	resourcesMu sync.Mutex
)

func reserveResource(id int) (*Resource, error) {
	resourcesMu.Lock()
	defer resourcesMu.Unlock()

	if _, ok := resources[id]; ok {
		return nil, fmt.Errorf("resource %d already reserved", id)
	}

	r := &Resource{id: id}
	resources[id] = r

	return r, nil
}

func releaseResource(id int) {
	resourcesMu.Lock()
	defer resourcesMu.Unlock()

	delete(resources, id)
}

func processResource(id int) error {
	r, err := reserveResource(id)
	if err != nil {
		return err
	}

	defer releaseResource(id)

	// Simulate processing the resource
	fmt.Printf("Processing resource %d\n", id)

	return nil
}

func main() {
	if err := processResource(1); err != nil {
		fmt.Println(err)
	}

	if err := processResource(1); err != nil {
		fmt.Println(err)
	}
}
