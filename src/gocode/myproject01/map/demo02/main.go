package main

import "fmt"

func main() {
	// //map 的操作，新增、删除、修改、查询
	// // 创建一个map
	// myMap := make(map[string]int)
	// // 添加键值对
	// myMap["apple"] = 5
	// myMap["banana"] = 10
	// myMap["orange"] = 15
	// // 查询键值对
	// fmt.Println("Initial map:", myMap)
	// // 修改键值对
	// myMap["banana"] = 20
	// fmt.Println("After modification:", myMap)

	// // 删除键值对
	// delete(myMap, "orange")
	// fmt.Println("After deletion:", myMap)
	// // 查询一个不存在的键
	// value, exists := myMap["banana"] // 查询键 "grape" 是否存在 ,value是对应的值，exists是一个布尔值，表示键是否存在
	// if exists {
	// 	fmt.Println("banana exists with value:", value)
	// } else {
	// 	fmt.Println("banana does not exist.")
	// }

	// //map长度
	// fmt.Println("Map length:", len(myMap))

	//map创建，key, value里再加map
	nestedMap := make(map[string]map[string]int)
	//赋值
	nestedMap["fruits"] = make(map[string]int)
	nestedMap["fruits"]["apple"] = 5
	nestedMap["fruits"]["banana"] = 10
	nestedMap["fruits"]["orange"] = 15
	fmt.Println("Nested map:", nestedMap)

	nestedMap["fruits2"] = make(map[string]int)
	nestedMap["fruits2"]["grape"] = 20
	nestedMap["fruits2"]["kiwi"] = 25
	nestedMap["fruits2"]["melon"] = 30
	fmt.Println("Nested map after adding fruits2:", nestedMap)

	// 遍历嵌套的map,nestedMap
	for category, items := range nestedMap {
		fmt.Println("Category:", category)
		for item, quantity := range items {
			fmt.Printf("  %s: %d\n", item, quantity)
		}
	}
}
