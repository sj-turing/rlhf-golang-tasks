import (
	"sync"
	"testing"
)

var globalVar int = 0
var wg sync.WaitGroup

func incrementGlobalVar(n int) {
	for i := 0; i < n; i++ {
		globalVar++
	}
	wg.Done()
}

func TestGarbageCollectionImpactConcurrency(t *testing.T) {
	const numGoroutines = 100
	const numIncrements = 10000

	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go incrementGlobalVar(numIncrements)
	}

	wg.Wait()

	expected := numGoroutines * numIncrements
	if globalVar != expected {
		t.Errorf("Expected %d, got %d", expected, globalVar)
	}
}
