package main

import "fmt"

func main() {
	// make函数创建切片
	slice := make([]int, 5, 11) // 创建一个长度为5，容量为11的切片
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[0] = 66
	slice[1] = 22
	slice[2] = 33
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v] = %v\n", i, slice[i])
	}

	fmt.Println("\n----------")

	//for range遍历切片
	for index, value := range slice {
		fmt.Printf("slice[%v] = %v\n", index, value)
	}

}
