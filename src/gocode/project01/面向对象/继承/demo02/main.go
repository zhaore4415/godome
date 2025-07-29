package main

import "fmt"

//结构体的组合模式
// Animal 结构体
type Animal struct {
	Name   string
	Age    int
	Weight float64
}

// Speak 方法
func (a *Animal) Speak() {
	fmt.Println("Animal speaks")
}

// ShowInfo 方法
func (b *Animal) ShowInfo() {
	fmt.Printf("Name: %s, Age: %d, Weight: %.2f\n", b.Name, b.Age, b.Weight)
}

// Dog 结构体，嵌入 Animal
type Dog struct {
	Animal
	b Animal  //这里把结构体当成变量
	int
}

// Speak 方法
func (d *Dog) Speak() {
	fmt.Println("Dog barks")
}

// ShowInfo 方法
func (d *Dog) ShowInfo() {
	fmt.Printf("~~~~~Name: %s, Age: %d, Weight: %.2f\n", d.Name, d.Age, d.Weight)
}
func main() {
	var a Animal
	a.Speak()
	a.ShowInfo()

	var d Dog
	d.Speak()
	d.ShowInfo()
	d.int = 10
	fmt.Println("Embedded int value:", d.int)
	// 结构体的组合模式
	d.b.Speak()
	d.b.ShowInfo()

}
