package main

import "fmt"

func main() {
	//updateSliceVal()
	deleteInForRange()
}

func updateSliceVal() {
	arr := make([]int, 10)

	// 问题：在range数组或切片时，是否可以通过v=1这种方式改变数组或切片的元素？
	// 答案：不可以，v只是临时变量，改变值只是对该变量有效，不能改变元素
	for i, v := range arr {
		v = 1
		fmt.Println("改变的值：", v)
		fmt.Println("切片内的值：", arr[i])
	}
}

func deleteInForRange() {
	m := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	for k := range m { // map 的遍历时无序的
		if k == 1 {
			fmt.Println("Delete key: ", k)
			// delete方法只能作用在map上
			delete(m, k)
			if _, ok := m[k]; ok {
				fmt.Println("Should not be ok!")
			}
		} else {
			fmt.Println("Remains key: ", k)
		}
	}
}
