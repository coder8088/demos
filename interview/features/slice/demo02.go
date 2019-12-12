package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	// 创建容量为5，长度为0的切片
	a := make([]int, 0, 5)
	fmt.Printf("%p\n", a)                                               // 0xc000080030
	fmt.Printf("%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)).Data) // 824634245168

	// 切片没有发生扩容 地址不变
	a = append(a, 1)
	fmt.Printf("%p\n", a)                                               // 0xc000080030
	fmt.Printf("%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)).Data) // 824634245168

	// 切片发生扩容之后 地址也发生变化
	a = append(a, []int{2, 3, 4, 5, 6}...)
	fmt.Printf("%p\n", a)                                               // 0xc00008c000
	fmt.Printf("%v\n", (*reflect.SliceHeader)(unsafe.Pointer(&a)).Data) // 824634294272
}
