package main

import (
	"fmt"
	"reflect"
)

type bs struct {
	name string
}

func main() {
	var (
		a int
		b bs
		c *bs
	)

	a = 10
	b = bs{name: "Jack"}
	c = &bs{name: "Ella"}

	rTypeA := reflect.TypeOf(a)
	rValA := reflect.ValueOf(a)

	fmt.Printf("a type: %v val: %v\n", rTypeA, rValA)

	rTypeB := reflect.TypeOf(b)
	rValB := reflect.ValueOf(b)

	fmt.Printf("b type: %v val: %v\n", rTypeB, rValB)

	rTypeC := reflect.TypeOf(c)
	rValC := reflect.ValueOf(c)

	fmt.Printf("c type: %[1]v val: %v\n", rTypeC, rValC)
}
