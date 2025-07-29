package main

import "fmt"

func test(num int) { //定义了一个test函数，没有参数，没有返回值
	fmt.Println(num)
}

func test02(num1 int, num2 float32, testFunc func(int)) { //定义了一个test02函数，没有参数，没有返回值
	fmt.Println("test02")
}
func main() {

	// var a int = 10
	// test(a) //调用test函数，传入参数a
	a := test
	fmt.Printf("a:%T\n,test:%T\n", a, test)              //打印a的类型
	a(10)                                                //调用test函数，传入参数10
	test02(10, 20.5, test)                               //调用test02函数，传入参数10,20,test函数
	test02(10, 20.5, a)                                  //调用test02函数，传入参数10,20,a函数
	test02(10, 20.5, func(num int) { fmt.Println(num) }) //调用test02函数，传入参数10,20,匿名函数
}
