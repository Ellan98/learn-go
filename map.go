package main

import "fmt"

// 无序的键值对集合  通过key快速检索数据  key 类似索引  指向数据的值Map是一种集合 所以可以向地带数组和切片那样迭代它
//Map是无序的 我们无法决定他返回顺序 这是因为Map是使用hash表来实现的 也是引用类型
//tips
//每次打印的map都会不一样 也是一种引用类型
//内置的len函数 同样适用于，map 返回map拥有的key的数量
//mao的key可以说所有可比较的类型 如布尔型，整数型   浮点型 复杂性等

//格式
// var map_variable map[key_data_type]value_data_type  声明变量 默认map 是 nil
//使用make ()创建map
//map_variabke = make(map[key_data_type]value_data_type)
func main() {
	rating := map[string]float32{"c": 5, "go": 4.5, "python": 4.5}
	fmt.Println(rating)
	EditMapData()
}

func EditMapData() {
	//如果不初始化map 就会创建一个nil map  nil map 不能用来存放键值对
	// var countryCapitalMap map[string]string
	countryCapitalMap := make(map[string]string)
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"
	countryCapitalMap["Japan"] = "Tokyo"
	countryCapitalMap["India"] = "New Delhi"

	for country := range countryCapitalMap {
		fmt.Println("Capital of ", country, "is", countryCapitalMap[country])
	}

	//map的长度  len()
	fmt.Println("countryCapitalMap的长度为", len(countryCapitalMap))

}
