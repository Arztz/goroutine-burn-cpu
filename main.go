package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Printf("Running with %d CPU cores\n", numCPU)
	runtime.GOMAXPROCS(8)
	var wg sync.WaitGroup
	wg.Add(numCPU)

	// Calculate the number of iterations needed to use approximately
	// 0.25 CPU core (assuming a CPU speed of 1 GHz)
	const iterationsPerCycle = 1000000
	targetCyclesPerSecond := 1000
	targetCyclesPerMillisecond := targetCyclesPerSecond / 1000
	targetCyclesPerMicrosecond := targetCyclesPerMillisecond / 1000
	targetCyclesPerNanosecond := targetCyclesPerMicrosecond / 1000
	iterationsPerSecond := 4 * targetCyclesPerNanosecond * 1000
	iterationsPerGoroutine := iterationsPerSecond / numCPU / targetCyclesPerSecond

	for i := 0; i < numCPU; i++ {
		go func() {
			for {
				// Burn CPU resources
				for j := 0; j < iterationsPerGoroutine; j++ {
				}
				runtime.Gosched()
			}
			wg.Done()
		}()
	}

	for {
		fmt.Printf("Number of Go routines: %d\n", runtime.NumGoroutine())
		runtime.Gosched()
	}

	wg.Wait()
}
