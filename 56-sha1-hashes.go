package main

import (
	"crypto/sha1"
	. "fmt"
)

func main() {
	s := "sha this string"
	hash := sha1.New()
	hash.Write([]byte(s))
	bs := hash.Sum(nil)
	Printf("%x\n", bs)
}
