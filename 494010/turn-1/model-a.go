package main

import (
	"io/ioutil"
	"log"
	"os"
)

func readFile1() (*os.File, error) {
	file, err := os.Open("file.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()
	// Use file
	//
	return file, nil
}

func readFile2() (*os.File, error) {
	file, err := os.Open("file.txt")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("Error closing file:", err)
		}
	}()
	return file, nil
}

// Use file
func useFile() ([]byte, error) {
	file, err := os.Open("file.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func main() {

	readFile1()
	readFile2()
	useFile()
}
