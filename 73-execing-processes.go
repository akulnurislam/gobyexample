package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"
)

func main() {
	binary, err := exec.LookPath("ls")
	if err != nil {
		panic(err)
	}
	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	err = syscall.Exec(binary, args, env)
	if err != nil {
		panic(err)
	}

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}
