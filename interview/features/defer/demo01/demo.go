package main

import "fmt"

func main() {
	// defer实现逆序输出
	a := []int{1, 2, 3, 4, 5}

	// 正常遍历
	for _, v := range a {
		fmt.Printf("%v\t", v)
	}

	fmt.Println()
	// defer逆序
	// defer不建议在for循环中使用，此处仅为示例
	for _, v := range a {
		defer func(v int) { fmt.Printf("%v\t", v) }(v)
	}
}
