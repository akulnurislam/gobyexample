package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() { messages <- "pang" }()
	go func() { messages <- "ping" }()
	go func() { messages <- "pong" }()

	// message := <-messages
	fmt.Println(<-messages)
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
