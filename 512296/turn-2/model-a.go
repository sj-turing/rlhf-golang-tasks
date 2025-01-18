import (
	"os"
	"runtime/pprof"
	"runtime"
	"fmt"
)

func main() {
	// Start CPU profiling
	f, err := os.Create("cpu.prof")
	if err != nil {
		fmt.Println("Could not create CPU profile: ", err)
		return
	}
	defer f.Close()
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Println("Could not start CPU profiling: ", err)
		return
	}
	defer pprof.StopCPUProfile()

	// Run your function(s) that you want to profile here
	runMatching()

	// Capture memory profile
	memProfile, err := os.Create("mem.prof")
	if err != nil {
		fmt.Println("Could not create memory profile: ", err)
		return
	}
	defer memProfile.Close()
	runtime.GC() // get up-to-date statistics
	if err := pprof.WriteHeapProfile(memProfile); err != nil {
		fmt.Println("Could not write memory profile: ", err)
		return
	}
}

func runMatching() {
	match := DNAMatch("AGCTAGCTAGCTAGCTAGCT", "AGCTAGCTAGCTAGCTAGCT")
	fmt.Println("Match:", match)
}
