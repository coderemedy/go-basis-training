package main

import "fmt"

type person struct {
	name string
	age  uint
}

func (p person) GoString() string {
	return fmt.Sprintf("person: {name: %s, age: %d}", p.name, p.age)
}

func main() {
	person := person{name: "John", age: 55}

	fmt.Printf("%#v\n", person)
}
