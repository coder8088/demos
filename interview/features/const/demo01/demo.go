package main

import "fmt"

type aa struct {
	name string
}

type bb func() string

type cc interface {
	say() string
}

// 非基础类型，不能定义为常量
//const aaS = aa{name: "Jack"}
//const bbF = bb
//const ccI = cc
//const aap = &aa{name:"Jack"}

func main() {
	const a = 100
	const b = 10
	const c = a / b

	fmt.Println(a, b, c)
}
