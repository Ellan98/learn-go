package main

// 访问数组元素
// float32 salary = balance[9]

import "fmt"

//数组的长度
// 通过将数组作为参数传递给len函数，可以获得数组的长度。
func main() {
	var n [10]int /* n 是一个长度为 10 的数组 */
	var i, j int

	/* 为数组 n 初始化元素 */
	for i = 0; i < 10; i++ {
		n[i] = i + 100 /* 设置元素为 i + 100 */
	}

	/* 输出每个数组元素的值 */
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}
	//数组的长度
	fmt.Println("length of a is", len(n))

	// for... range...
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	// 如果需要值并希望忽略索引，那么可以通过使用_ blank标识符替换索引来实现这一点。
	// for _, v := range a { //ignores index
	// }

	fmt.Println("\nsum of all elements of a", sum)
}
