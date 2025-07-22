package main

import "fmt"

//全局变量

func main() {
	// 声明变量
	var a int = 10
	var b int = 20
	var sum int = a + b
	// 计算和
	//	sum = a + b
	// 打印结果
	fmt.Println("Sum of", a, "and", b, "is", sum)
	fmt.Println("Hello, world!")

	// 第二种声明变量
	var aa int
	fmt.Println("请输入一个整数：", aa)

	//第三种声明变量
	var aaa = 20
	fmt.Println("等于：", aaa)

	//第四种声明变量,如果不指定类型，Go 会根据值自动推断类型，必须加上:
	sex := "100"
	fmt.Println(sex)
	fmt.Println("-----------------------------------------------------")
	//声明多个变量
	var c, d int = 10, 20
	fmt.Println("c =", c, "d =", d)

	var e, f, g int
	fmt.Println("e =", e, "f =", f, "g =", g)

	h, j := 100, 200
	fmt.Println("h =", h, "j =", j) // 声明多个变量并同时赋值

	//常量

}
