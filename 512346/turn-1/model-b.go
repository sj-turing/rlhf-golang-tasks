package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// LogFileData struct contains the file path and last modification time
type LogFileData struct {
	path    string
	lastMod time.Time
}

// Chan hold the files with its last modification time
var lfdMap = make(map[string]LogFileData)
var mutex = &sync.Mutex{}

// Function to watch the changes in the folder
func watchFolder(path string) {
	for {
		mutex.Lock()
		for filePath, lfd := range lfdMap {
			//Get the file information
			fi, err := os.Stat(filePath)
			if err != nil {
				// If the file doesn't exist, remove it from the map
				if os.IsNotExist(err) {
					delete(lfdMap, filePath)
				} else {
					log.Println(err)
				}
			} else if fi.ModTime() != lfd.lastMod {
				//If the file is modified, call the function to process the modified file
				ProcessModifiedFile(filePath)
			}
		}
		mutex.Unlock()
		time.Sleep(time.Second * 1) // Check every second
	}
}

// Function to process the modified file
func ProcessModifiedFile(filePath string) {
	fmt.Println("File changed:", filePath)
	// Process the modified file and update the visualization in LOKI
}

// Function to add a new file in the map
func addNewFile(filePath string) {
	mutex.Lock()
	fi, _ := os.Stat(filePath)
	lfdMap[filePath] = LogFileData{path: filePath, lastMod: fi.ModTime()}
	mutex.Unlock()
}

func main() {
	go watchFolder("./logs") // Replace './logs' with the actual folder path

	// Add initial log files to the map
	files, err := filepath.Glob("./logs/*.log")
	if err != nil {
		log.Fatal(err)
	}
	for _, filePath := range files {
		addNewFile(filePath)
	}
	// Infinite loop to keep the application running
	for {
		time.Sleep(time.Minute * 1) //Add files every minute
	}
}
