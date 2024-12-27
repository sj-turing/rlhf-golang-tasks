import (
	"runtime"
)

func TestGCStats(t *testing.T) {
	var gcStats runtime.GCStats
	// Before starting work
	runtime.ReadGCStats(&gcStats)
	for i := 0; i < 10000; i++ {
		globalVar++
		if i%1000 == 0 {
			// Simulate some work
			time.Sleep(100 * time.Millisecond)
		}
	}
	// After work
	runtime.ReadGCStats(&gcStats)
	t.Log("GC Stats:", gcStats)
}
