package main  
import (  
    "context"
    "fmt"
    "math/rand"
    "sync"
    "time"
)
func fetchData(ctx context.Context, url string) error {
    // Simulate HTTP request duration with a random delay
    time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
    fmt.Println("Fetched data from:", url)
    return nil
}
func main() {
    urls := []string{"url1", "url2", "url3", "url4", "url5"}
    const numGoroutines = 3

    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    wg := sync.WaitGroup{}
    wg.Add(numGoroutines)

    dataCh := make(chan error, len(urls))

    go func() {
        for err := range dataCh {
            if err != nil {
                fmt.Println("Error:", err)
                cancel() // Cancel context if there's an error
                break
            }
        }
        wg.Done()
    }()

    for i := 0; i < numGoroutines; i++ {
        go func() {
            for _, url := range urls {
                select {
                case <-ctx.Done():
                    return // Return if context is canceled
                default:
                    err := fetchData(ctx, url)
