package main

import (
	"fmt"
)

type adder = struct {
}

/*func (a adder) Add(aa int, b ...int) int {
	for _, v := range b {
		aa += v
	}
	return aa
}*/

func main() {
	a := adder{}
	fmt.Printf("a type: %T\n", a)
}
