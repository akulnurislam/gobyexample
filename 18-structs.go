package main

import "fmt"

type person struct {
	name string
	age  int
}

func NewPerson(name string) *person {
	p := person{name: name}
	p.age = 66
	return &p
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{age: 30, name: "Alice"})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{"Ann", 40})
	fmt.Println(NewPerson("Akul"))

	sean := person{name: "Sean", age: 50}
	fmt.Println(sean.name)
	fmt.Println(sean.age)

	sp := &sean
	fmt.Println(sp.age)
	sp.age = 70

	fmt.Println(sean)
}
