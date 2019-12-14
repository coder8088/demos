package main

import (
	"fmt"
	"reflect"
)

type person struct {
	name string
}

func main() {
	var p1 *person
	var p2 *person

	register(&p1)
	register(&p2)

	fmt.Println(p1)
	fmt.Println(p2)
}

func register(target interface{}) {
	p := &person{name: "jack"}

	tv := reflect.ValueOf(target)
	pv := reflect.ValueOf(p)
	fmt.Printf("target type:%v\n", tv.Type())
	fmt.Printf("source type:%v\n", pv.Type())

	tv.Elem().Set(pv)
}
