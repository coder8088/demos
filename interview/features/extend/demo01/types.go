package demo01

type Person struct {
	Name string
	age  int
}

type Adder = interface {
	Add(a int, b ...int) int
}

type AddFunc = func(a int, b ...int) int

type Person2 = struct {
	Name string
	Age  int
}
