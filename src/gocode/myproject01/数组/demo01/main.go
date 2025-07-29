package main

import "fmt"

func main() {
	testArray()
}

//五种数组初始化方式
func testArray() {
	// 1. 使用 var 声明数组
	var arr1 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("使用 var 声明数组并初始化：", arr1)

	// 2. 使用 var 声明数组并指定长度
	var arr2 [5]int
	arr2[0] = 1
	arr2[1] = 2
	arr2[2] = 3
	arr2[3] = 4
	arr2[4] = 5
	fmt.Println("使用 var 声明数组：", arr2)
	// 3. 使用数组字面量初始化
	var arr3 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("使用数组字面量初始化：", arr3)

	// 4. 使用省略号初始化
	var arr4 = [...]int{1, 2, 3, 4, 5}
	fmt.Println("使用省略号初始化：", arr4)
	// 指定索引位置初始化
	var arr5 = [...]int{2: 66, 0: 33, 1: 99, 3: 88, 4: 77}
	fmt.Println("使用省略号初始化：", arr5)

	// 5. 使用 make 函数创建数组
	// 注意：Go语言中没有 make 函数用于创建数组，make 函数通常用于创建切片、映射和通道。
	// 但是可以使用 make 函数创建切片，然后将其转换为数组
	var arr6 = make([]int, 5) // 创建一个长度为5的切片
	arr6[0] = 1
	arr6[1] = 2
	arr6[2] = 3
	arr6[3] = 4
	arr6[4] = 5
	var arr7 = [...]int{arr6[0], arr6[1], arr6[2], arr6[3], arr6[4]} // 转换为数组
	fmt.Println("使用 make 函数创建数组：", arr7)

	//循环数组
	var arr [5]int
	for i := 0; i < 5; i++ {
		arr[i] = i + 1
	}
	fmt.Println("数组内容：", arr)

	//第arr7 用range 循环遍历数组
	for index, value := range arr7 {
		fmt.Printf("索引：%d, 值：%d\n", index, value)
	}

}
