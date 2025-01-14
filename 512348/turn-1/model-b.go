package main
import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "sync"
)

// Platform-specific directory separator
const dirSep = string(os.PathSeparator)

// batchDir creates a directory batch operation struct.
type batchDir struct {
    dirs []string
    wg   *sync.WaitGroup
}

// newBatchDir creates a new batchDir struct.
func newBatchDir(dirs []string, wg *sync.WaitGroup) *batchDir {
    return &batchDir{dirs: dirs, wg: wg}
}

// createDirs creates directories in a batch using goroutines.
func (bd *batchDir) createDirs() {
    defer bd.wg.Done()
    for _, dir := range bd.dirs {
        err := os.MkdirAll(dir, 0755)
        if err != nil {
            log.Printf("Error creating directory '%s': %v", dir, err)
        }
    }
}

// deleteDirs deletes directories in a batch using goroutines.
func (bd *batchDir) deleteDirs() {
    defer bd.wg.Done()
    for _, dir := range bd.dirs {
        err := os.RemoveAll(dir)
        if err != nil {
            log.Printf("Error deleting directory '%s': %v", dir, err)
        }
    }
}

// batchProcess batches the directory creation or deletion operations based on the provided mode.
func batchProcess(dirs []string, mode string, batchSize int) {
    var wg sync.WaitGroup
    for i := 0; i < len(dirs); i += batchSize {
        end := i + batchSize
        if end > len(dirs) {
            end = len(dirs)
        }
        bd := newBatchDir(dirs[i:end], &wg)
        wg.Add(1)
        if mode == "create" {
            go bd.createDirs()
        } else if mode == "delete" {
            go bd.deleteDirs()
        } else {
            log.Printf("Invalid mode '%s'. Mode must be 'create' or 'delete'.", mode)
            return
        }
    }
    wg.Wait()
}

// main function of the bulk directory management program.
func main() {
    if len(os.Args) != 4 {
        fmt.Println("Usage: go run main.go <mode> <root_dir> <batch_size>")
        fmt.Println("Example: go run main.go create /tmp/bulk_dirs 100")
        fmt.Println("Example: go run main.go delete /tmp/bulk_dirs 100")
        return
    }

    mode := os.Args[1]
    rootDir := os.Args[2]
    batchSizeStr := os.Args[3]

    batchSize, err := strconv.Atoi(batchSizeStr)
    if err != nil {
        log.Fatalf("Invalid batch size '%s': %v", batchSizeStr, err)
    }

    // Create or delete directories in bulk based on the provided mode
    if mode == "create" {
        // Create a directory hierarchy
        dirs := make([]string, 0)
        err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
            if err != nil {
                return err
            }
            if info.IsDir() {
                dirs = append(dirs, path)
            }
            return nil
        })
        if err != nil {
            log.Fatalf("Error walking directory '%s': %v", rootDir, err)
        }
        batchProcess(dirs, "create", batchSize)
    } else if mode == "delete" {
