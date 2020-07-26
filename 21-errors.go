package main

import (
	"errors"
	"fmt"
)

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

type customError struct {
	argument int
	problem  string
}

func (e customError) Error() string {
	return fmt.Sprintf("%d - %s", e.argument, e.problem)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &customError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main() {
	a, err := f1(42)
	if err != nil {
		fmt.Println(err)
	}

	b, err := f2(42)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a)
	fmt.Println(b)

	fmt.Println()

	_, e := f2(42)
	if ce, ok := e.(*customError); ok {
		fmt.Println("extract custom error")
		fmt.Println(ce.argument)
		fmt.Println(ce.problem)
	}
}
