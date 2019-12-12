package main

import "fmt"

type Person1 struct {
	Name string
	Age  int
}

func main() {
	p1 := &Person1{"jack", 19}
	p2 := &Person1{"Ella", 18}

	fmt.Printf("%p\t%#v\n", p1, p1)
	fmt.Printf("%p\t%#v\n", p2, p2)

	*p1 = *p2

	fmt.Printf("%p\t%#v\n", p1, p1)
	fmt.Printf("%p\t%#v\n", p2, p2)

}
