package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	fmt.Println("before:", strs)
	sort.Strings(strs)
	fmt.Println("after :", strs)
	fmt.Println()

	ints := []int{7, 2, 4}
	fmt.Println("before:", ints)
	fmt.Println("sorted:", sort.IntsAreSorted(ints))
	sort.Ints(ints)
	fmt.Println("after :", ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted:", s)
}
