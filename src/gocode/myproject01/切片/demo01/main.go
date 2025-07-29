package main

import "fmt"

func main() {

	// 直接定义切片
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	c := []int{7, 8, 9}
	// 打印切片
	fmt.Println(a, b, c)

	// 使用append函数添加元素
	c = append(c, 4)
	// 打印切片
	fmt.Println(c)

	//定义数组并从数组创建切片
	d := [8]int{1, 2, 3, 5, 9, 8, 10, 4} // 定义一个数组
	slice1 := d[1:3]                     // 切片从索引1到索引3（不包括索引3）   -[1,3)  是左闭右开区间
	slice2 := d[2:4]                     // 切片从索引2到索引4（不包括索引4）   -[2,4)  是左闭右开区间
	slice3 := d[1:4]                     // 切片从索引1到索引4（不包括索引4）   -[1,4)  是左闭右开区间
	fmt.Println(slice1, slice2, slice3)

	len := len(slice1) // 获取长度
	cap := cap(slice1) // 获取容量
	fmt.Println(len, cap)
}
