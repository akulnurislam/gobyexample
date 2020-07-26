package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

func main() {
	spawn("date")
	spawn("cal")

	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Printf("> %s\n", "grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("> %s\n", "ls -a -l -h")
	fmt.Println(string(lsOut))

	ticker := time.NewTicker(time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Println(time.Now())
		}
	}
}

func spawn(command string) {
	dateCmd := exec.Command(command)
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Printf("> %s\n", command)
	fmt.Println(string(dateOut))
}
