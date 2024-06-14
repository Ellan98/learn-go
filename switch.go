package main

import "fmt"

// 程序的流程控制结构一共有三种：顺序结构，选择结构，循环结构。
func main() {
	/* 定义局部变量 */
	var grade string = "B"
	var marks int = 90

	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C" //case 后可以由多个数值
	default:
		grade = "D"
	}

	switch {
	case grade == "A":
		fmt.Printf("优秀!\n")
	case grade == "B", grade == "C":
		fmt.Printf("良好\n")
	case grade == "D":
		fmt.Printf("及格\n")
	case grade == "F":
		fmt.Printf("不及格\n")
	default:
		fmt.Printf("差\n")
	}
	fmt.Printf("你的等级是 %s\n", grade)
}

// switch var1 {
// case val1:
// 		...
// case val2:
// 		...
// default:
// 		...
// }
// Go里面switch默认相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码。

// 类型不被局限于常量或整数，但必须是相同的类型；或者最终结果为相同类型的表达式。 可以同时测试多个可能符合条件的值，使用逗号分割它们，例如：case val1, val2, val3。
// 如需贯通后续的case，就添加fallthrough

// Tips
// case后的常量值不能重复
// case后可以有多个常量值
// fallthrough应该是某个case的最后一行。如果它出现在中间的某个地方，编译器就会抛出错误。
