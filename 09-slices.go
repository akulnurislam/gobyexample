package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	fmt.Println("len:", len(c))
	fmt.Println("cap:", cap(c))

	l := c[2:5]
	fmt.Println("sl1:", l)
	fmt.Println("len:", len(l))
	fmt.Println("cap:", cap(l))

	l = c[:5]
	fmt.Println("sl2:", l)
	fmt.Println("len:", len(l))
	fmt.Println("cap:", cap(l))

	l = c[2:]
	fmt.Println("sl3:", l)
	fmt.Println("len:", len(l))
	fmt.Println("cap:", cap(l))

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d :", twoD)
}
