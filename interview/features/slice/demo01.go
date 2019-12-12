package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	a := []int{1, 3, 4, 5, 3}
	b := a[1:3]
	c := a[:]

	// a和b中Data指针的值变得不一样
	fmt.Printf("%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)).Data) // 824634166960
	fmt.Printf("%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&b)).Data) // 824634166968
	// a和c中Data指针的值一样
	fmt.Printf("%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&c)).Data) // 824634166960
}
