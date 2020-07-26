package main

import (
	"fmt"
	"time"
)

func fact(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	start := time.Now()
	result := fact(50)
	elapsed := time.Since(start)
	fmt.Println(result, elapsed)
}
