package main

import "fmt"

//方法和函数的区别
// 区别1 定义方式

type Person struct {
	Name string
}

// 定义方法  , 一这要有类型
func (s Person) test() {
	fmt.Println("This is a method of Person.")
}

// 定义函数
func test(s Person) {
	fmt.Println("This is a function.", s.Name)
}

func main() {
	p := Person{Name: "Alice"}
	p.test() // 调用方法

	test(p) // 调用函数
}
