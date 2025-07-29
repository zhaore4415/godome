package main

//改变值类型内部的值
// 如果想改变值类型内部的值，可以用指针做引用传递，如下，&num  就是把地址传进去，再修改地址对应的值，
// 最后输出时，num变成变成了30

import "fmt"

//参数的类型为指针
func test(num *int) {
	//对地址对应的变更进行改变值
	*num = 30
}

func main() {

	var num int = 10
	fmt.Println(num) //输出10
	//	fmt.Println(&num)  //输出 内存引用地址   &num 就是地址
	test(&num)
	fmt.Println(num) //输出30
}
