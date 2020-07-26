package main

import (
	. "fmt"
	"math/rand"
	"time"
)

func main() {
	Println(rand.Intn(100))
	Println(rand.Intn(100))
	Println()

	Println(rand.Float64())
	Println()

	// 5.0 <= f' < 10.0
	Println((rand.Float64() * 5) + 5)
	Println((rand.Float64() * 5) + 5)
	Println()

	source1 := rand.NewSource(time.Now().UnixNano())
	random1 := rand.New(source1)
	Printf("%d,%d\n", random1.Intn(100), random1.Intn(100))

	source2 := rand.NewSource(42)
	random2 := rand.New(source2)
	Printf("%d,%d\n", random2.Intn(100), random2.Intn(100))

	source3 := rand.NewSource(42)
	random3 := rand.New(source3)
	Printf("%d,%d\n", random3.Intn(100), random3.Intn(100))
}
