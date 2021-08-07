package main

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func main() {
	findType := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		case Person:
			fmt.Println("Person")
		default:
			fmt.Printf("don't know type: %T\n", t)
		}
	}
	findType(true)
	findType(1)
	findType("hey")

	p := Person{name: "Bob", age: 3}
	findType(p)
}
