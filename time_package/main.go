package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("当前时间为:%v\n", timeDemo())
	now := time.Now()

	fmt.Println(now.Format("2006-01-02 15:04:05"))
}

// timeDemo 时间对象的年月日时分秒
func timeDemo() string {
	now := time.Now() // 获取当前时间
	fmt.Printf("current time:%v\n", now)

	year := now.Year()     // 年
	month := now.Month()   // 月
	day := now.Day()       // 日
	hour := now.Hour()     // 小时
	minute := now.Minute() // 分钟
	second := now.Second() // 秒
	fmt.Println(year, month, day, hour, minute, second)
	return fmt.Sprint(year, month, day, hour, minute, second)
}
