package main

import "fmt"

func main() {
	arr := [3]int{1, 2, 3}         // [1, 2, 3]
	brr := [...]int{1, 2, 3}       // [1, 2, 3]
	crr := [...]int{2: 3, 1: 2}    // [0, 2, 3]
	drr := [...]int{1, 2, 4: 5, 6} // [1, 2, 0, 0, 5, 6]

	// 数组变量的类型就是数组类型本身，而不是指针，数组长度是数组类型的一部分
	fmt.Printf("%T\n", arr) // [3]int
	fmt.Printf("%v\n", arr)
	fmt.Printf("%v\n", brr)
	fmt.Printf("%v\n", crr)
	fmt.Printf("%v\n", drr)

	slice := []int{1, 2, 3}

	fmt.Printf("%T\n", slice) // []int
}
