package pkg

import "fmt"

func init() {
	fmt.Println("执行pkg02.go中的init函数...")

	// init中启动的goroutine只会在进入main.main函数后才可能被执行
	go func() {
		fmt.Println("pkg02.go的init方法中调用的goroutine...")
	}()
}

func B() {
	fmt.Println("执行pkg02.go中的B函数...")
}
