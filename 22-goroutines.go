package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("%s -> %d\n", from, i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(message string) {
		fmt.Println(message)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("done")
}
