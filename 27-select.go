package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case message1 := <-c1:
			fmt.Println("received", message1)
		case message2 := <-c2:
			fmt.Println("received", message2)
		}
	}
}