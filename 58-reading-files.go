package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dat, err := ioutil.ReadFile("/tmp/dat")
	check(err)
	Print(string(dat))

	Println()

	f, err := os.Open("/tmp/dat")
	check(err)
	defer f.Close()

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	Println("b1", b1)
	Println("n1", n1)
	Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	Println()

	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	Println("o2", o2)
	Println("b2", b2)
	Println("n2", n2)
	Printf("%d bytes @ %d: ", n2, o2)
	Printf("%v\n", string(b2[:n2]))

	Println()

	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	Println("o3", o3)
	Println("b3", b3)
	Println("n3", n3)
	Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	_, err = f.Seek(0, 0)
	check(err)

	Println()

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	Printf("5 bytes: %s\n", string(b4))
}
