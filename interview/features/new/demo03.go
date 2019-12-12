package main

import "fmt"

func main() {
	a := a()
	fmt.Println(a)
}

func a() int {
	x := new(int) // 只是创建了一个int类型的指针，这个指针到底指向哪里是go在运行时帮我们搞定
	return *x
}
