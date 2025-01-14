package main  
import (  
    "fmt"
    "log"
    "os"
    "strconv"
    "sync"
)

const dirSep = string(os.PathSeparator)

// createDir creates a single directory and sends an error (if any) down the error channel.
func createDir(dir string, errc chan<- error) {  
    err := os.MkdirAll(dir, 0755)  
    if err != nil {  
        errc <- err  
        return  
    }  
    errc <- nil  
}

// deleteDir deletes a single directory and sends an error (if any) down the error channel.
func deleteDir(dir string, errc chan<- error) {  
    err := os.RemoveAll(dir)  
    if err != nil {  
        errc <- err  
        return  
    }  
    errc <- nil  
}

// batchProcess orchestrates the creation or deletion of directories in batches, handling errors as they occur.
func batchProcess(mode string, batchSize int, dirs []string) {  
    numBatches := len(dirs) / batchSize  
    if len(dirs)%batchSize > 0 {  
        numBatches++  
    }  

    // Create a channel to collect errors from goroutines.
    errc := make(chan error, batchSize)  

    // Ensure goroutines finish before main returns.
    var wg sync.WaitGroup  
    wg.Add(numBatches)  

    for start := 0; start < len(dirs); start += batchSize {  
        end := start + batchSize  
        if end > len(dirs) {  
            end = len(dirs)  
        }  

        go func(start, end int) {  
            defer wg.Done()  
            for _, dir := range dirs[start:end] {  
                switch mode {  
                case "create":  
                    createDir(dir, errc)  
                case "delete":  
                    deleteDir(dir, errc)  
                }  
            }  
        }(start, end)  
    }  

    // Wait for all goroutines to complete and collect errors.
    wg.Wait()  
    close(errc)  

    // Process errors, if any.
    for err := range errc {  
        if err != nil {  
            log.Printf("Error occurred: %v", err)  
        }  
    }  
}  

func main() {  
    // Check for correct usage.
    if len(os.Args) < 4 || len(os.Args) > 5 {  
        fmt.Println("Usage: go run main.go <mode> <root_dir> <batch_size> <directory_count>")  
        fmt.Println("Example: go run main.go create ./bulk_dirs 10 100")  
        fmt.Println("Example: go run main.go delete ./bulk_dirs 10")  
        return  
    }  
    // Parse command-line arguments.
    mode := os.Args[1]  
    rootDir := os.Args[2]  
    batchSizeStr := os.Args[3]  
    batchSize, err := strconv.Atoi(batchSizeStr)  
    if err != nil {  
        log.Fatalf("Invalid batch size '%s': %v", batchSizeStr, err)  
    }  
    
    switch mode {  
    case "create":  
        if len(os.Args) != 5 {  
            fmt.Println("Error: When creating directories, please specify the 'directory_count'.")
            return
        }
        directoryCountStr := os.Args[4]  
        directoryCount, err := strconv.Atoi(directoryCountStr)  
        if err != nil {  
            log.Fatalf("Invalid directory size '%s': %v", directoryCountStr, err)  
