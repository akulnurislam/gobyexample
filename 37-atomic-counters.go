package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var count uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {
				atomic.AddUint64(&count, 1)
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("count:", count)
}
