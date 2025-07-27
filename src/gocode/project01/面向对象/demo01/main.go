package main

import "fmt"

//结构体
type Person struct {
	Name string
	Age  int
}

// type Animal struct {
// 	Name    string // 动物的名字
// 	Species string // 动物的种类
// 	Age     int    // 动物的年龄
// }

func main() {
	y := Person{}
	var x Person      // 声明一个Person类型的变量x
	fmt.Println(y, x) //输出为{ 0} { 0}，因为结构体的字段都被初始化为零值

	p := Person{Name: "Alice", Age: 30}
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
	// 访问结构体字段
	p.Name = "Bob" // 修改字段值
	p.Age = 25     // 修改字段值

	fmt.Println("Updated Name:", p.Name)
	fmt.Println("Updated Age:", p.Age)
	// 结构体作为函数参数
	printPerson(p) // 传递结构体实例

	// // 结构体作为返回值
	// animal := createAnimal("Lion", "Panthera leo", 5)

	// fmt.Println("Animal:", animal)

}

func printPerson(p Person) {
	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)
}
