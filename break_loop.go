package main

import "fmt"

// break：跳出循环体。break语句用于在结束其正常执行之前突然终止for循环
// continue语句 跳出一次循环。continue语句用于跳过for循环的当前迭代。
func main() {
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
}
