package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {
	const JOBS int = 5

	jobs := make(chan int, JOBS)
	results := make(chan int, JOBS)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= JOBS; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a < JOBS; a++ {
		result := <-results
		fmt.Println("result:", result)
	}
}
