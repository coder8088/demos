package main

import (
	"fmt"
	"time"
)

// 测试goroutine中的panic怎么传播
func main() {
	fmt.Println("main run...")
	// 这里进行异常捕获不生效
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("main catch exception: ", r)
		}
	}()

	// goroutine 中的panic没有被recover会传播到main中
	go func() {
		// goroutine内捕获异常，panic就不会传播到main中
		/*defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()*/
		fmt.Println("goroutine run...")
		time.Sleep(2 * time.Second)
		panic("发生异常。。。")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("main exit")
}
