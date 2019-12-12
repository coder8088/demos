package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := new(Person)
	fmt.Printf("%T\t%#v", p, p)
}
