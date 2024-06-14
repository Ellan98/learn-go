package main

import (
	"fmt"
	"reflect"
)

var name string = "hello 字符串"

func main() {
	v := reflect.ValueOf(name)
	if v.Kind() == reflect.String {
		fmt.Printf("变量类型为%T\n 其值为%v", name, name)
	}
}

// 需要注意的是，num的定义在if里，那么只能够在该if..else语句块中使用，否则编译器会报错的。
// if 布尔表达式 {
// 	/* 在布尔表达式为 true 时执行 */
// }
// if 布尔表达式 {
// 	/* 在布尔表达式为 true 时执行 */
// } else {
//  /* 在布尔表达式为 false 时执行 */
// }
// if 布尔表达式1 {
// 	/* 在布尔表达式1为 true 时执行 */
// } else if 布尔表达式2{
// 	/* 在布尔表达式1为 false ,布尔表达式2为true时执行 */
// } else{
// 	/* 在上面两个布尔表达式都为false时，执行*/
// }
