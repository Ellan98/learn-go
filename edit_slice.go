package main

import "fmt"

//修改切片
//slice 只是底层数组的一个表示 对slice所做的任何修改都将反映在底层数组中
func main() {
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	// for i := range dslice {
	// 	dslice[i]++
	// }
	for i := 0; i < len(dslice); i++ {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}
