package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ch := make(chan int, 64) // 64容量的成果队列

	go produce(3, ch) // 生成3的倍数成果
	go produce(5, ch) // 生成5的倍数成果

	go consume(ch) // 消费成果
	go consume(ch)

	// Ctrl+C退出
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	fmt.Printf("quit (%v)\n", <-sigCh)
}

func produce(factor int, out chan<- int) {
	for i := 0; ; i++ {
		time.Sleep(2 * time.Second)
		out <- i * factor
	}
}

func consume(in <-chan int) {
	for v := range in {
		fmt.Println("consume: ", v)
	}
}
