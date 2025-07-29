package main

import "fmt"

//封装 相关的示例
type Person struct {
	Name string // 姓名
	Age  int    // 年龄
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	// 访问结构体字段
	p.Name = "Bob" // 修改字段值
	p.Age = 25     // 修改字段值

	printPerson(p) // 传递结构体实例
}

func printPerson(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
