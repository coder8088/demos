package main

import "fmt"

func main() {
	const s = "abc阿訇"
	for _, c := range s {
		fmt.Printf("%c\n", c) // a b c 阿 訇
	}

	// 对字符串进行截取操作，返回的结果还是string类型
	a := s[:2]
	fmt.Printf("%T\n", a) // string
}
