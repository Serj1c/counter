package counter

import (
	"runtime"
	"sync"
	"testing"
)

func BenchmarkCounter(b *testing.B) {
	wg := sync.WaitGroup{}
	numCPU := runtime.NumCPU()
	wg.Add(numCPU)

	counter := ShardedAtomicCounter{}
	for i := 0; i < numCPU; i++ {
		go func(idx int) {
			defer wg.Done()
			for j := 0; j < b.N; j++ {
				counter.Increment(idx)
			}
		}(i)
	}

	wg.Wait()
}
