package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4, 5}
	var b = &a

	fmt.Printf("a: %T\tb: %T\n", a, b)

	fmt.Printf("%v", b[2]) // 数组指针也可以像数组本身一样取值

	// 数组指针也可以迭代数组的元素
	for i, t := range b {
		fmt.Printf("%d: %d\n", i, t)
	}

	// 一个数组变量被赋值或者被传递时，实际上会复制整个数组
	var c = a
	var d = c
	fmt.Printf("a addr: %p\n", &a) // 0xc000018120
	fmt.Printf("c addr: %p\n", &c) // 0xc00009c000
	fmt.Printf("d addr: %p\n", &d) // 0xc00009c030
}
