package main

import (
	"fmt"
	"math"
)

type react struct {
	width, height int
}

func (r *react) area() int {
	return r.width * r.height
}

func (r *react) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := react{width: 10, height: 5}

	fmt.Println("area :", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r

	fmt.Println("area :", rp.area())
	fmt.Println("perim:", rp.perim())
}
