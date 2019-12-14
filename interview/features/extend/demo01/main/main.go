package main

import (
	"fmt"
	"github.com/tedux/demos/interview/features/extend/demo01"
)

type Student struct {
	demo01.Person
	class string
}

type Teacher struct {
	*demo01.Person
	lesson string
}

func main() {
	// 无论是机构体组合 还是指针类型组合，都要传入具体的匿名字段值
	stu := Student{demo01.Person{Name: "jack"}, "class1"}
	t := Teacher{&demo01.Person{Name: "Ella"}, "math"}

	//a := t.Person.age 不能访问私有的成员 编译不通过 -> Unexported field 'age' usage

	fmt.Println(stu.Name, t.Name)
}
