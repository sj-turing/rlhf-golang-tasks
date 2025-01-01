package main  
import (  
    "fmt"
    "sync"
    "time"
)
func main() {
    myMap := new(sync.Map)
    var wg sync.WaitGroup
    numGoroutines := 1000
    numOps := 10000

    wg.Add(numGoroutines)
    for i := 0; i < numGoroutines; i++ {
        go func() {
            for j := 0; j < numOps; j++ {
                myMap.Store(j, j)
            }
            wg.Done()
        }()
    }

    start := time.Now()
    wg.Wait()
    end := time.Now()

    var count int
    myMap.Range(func(_, _ interface{}) bool {
        count++
        return true
    })

    fmt.Printf("Execution time: %v\n", end.Sub(start))
