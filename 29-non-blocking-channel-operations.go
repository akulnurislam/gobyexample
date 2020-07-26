package main

import "fmt"

func main() {
	messages := make(chan string)

	select {
	case message := <-messages:
		fmt.Println("received message", message)
	default:
		fmt.Println("no message received")
	}

	message := "hi..."
	select {
	case messages <- message:
		fmt.Println("sent message", message)
	default:
		fmt.Println("no message sent")
	}

	signals := make(chan bool)
	select {
	case message := <-messages:
		fmt.Println("received message", message)
	case signal := <-signals:
		fmt.Println("received signal", signal)
	default:
		fmt.Println("no activity")
	}
}
