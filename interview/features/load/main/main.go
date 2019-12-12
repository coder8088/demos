package main

import (
	"fmt"
	"github.com/tedux/demos/interview/features/load/pkg"
	"time"
)

func main() {
	fmt.Println("main.main开始执行...")
	pkg.A()

	// 睡眠1秒，检查goroutine是否执行
	time.Sleep(time.Second)
}
