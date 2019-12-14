package main

// 交替打印数字字母
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	words := "abcdef"

	numCh := make(chan bool)
	alCh := make(chan bool, 1)
	done := make(chan int)

	go func() {
		var cur = 0
		for ; cur < len(nums); cur += 2 {
			<-alCh
			print(nums[cur])
			print(nums[cur+1])
			numCh <- true
		}
	}()

	go func() {
		var cur = 0
		for ; cur < len(words); cur += 2 {
			<-numCh
			print(string(words[cur]))
			print(string(words[cur+1]))
			alCh <- true
		}
		done <- 1
	}()

	alCh <- true
	<-done
}
