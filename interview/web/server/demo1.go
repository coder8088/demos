package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 创建简单的服务，回复hello word和时间
func main() {
	fmt.Println("Please visit http://127.0.0.1:12345")
	// HandleFunc函数是将pattern和对应的handler注册到ServerMux对象中
	// 访问 http://127.0.0.1:12345 或者 http://127.0.0.1:12345/cn* 都会执行这个handler
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprintf("你好，世界！-- Time: %s", time.Now())
		fmt.Fprintf(w, "%s\n", s)
		log.Printf("%v\n", s)
	})

	// 访问 http://127.0.0.1:12345/en 会执行这个handler
	http.HandleFunc("/en", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprintf("Hello, world! -- Time: %s", time.Now())
		fmt.Fprintf(w, "%s\n", s)
		log.Printf("%v\n", s)
	})

	// 识别不到
	http.HandleFunc("enn", func(w http.ResponseWriter, req *http.Request) {
		s := fmt.Sprintf("Hello, world! -- Time: %s", time.Now())
		fmt.Fprintf(w, "%s\n", s)
		log.Printf("%v\n", s)
	})

	if err := http.ListenAndServe(":12345", nil); err != nil {
		panic(err)
	}
}
