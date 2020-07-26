package main

import (
	. "fmt"
	"strconv"
)

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)
	Println(f)

	i, _ := strconv.ParseInt("123", 0, 64)
	Println(i)

	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	Println(d)

	u, _ := strconv.ParseUint("789", 0, 64)
	Println(u)

	k, _ := strconv.Atoi("135")
	Println(k)

	_, e := strconv.Atoi("wat")
	Println(e)
}
