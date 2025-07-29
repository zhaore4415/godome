package main

import "fmt"

//动物有名字
type Animal struct {
	Name   string
	Age    int
	Weight float64
}

func (a *Animal) Speak() {
	fmt.Println("Animal speaks")
}
func (b *Animal) ShowInfo() {
	fmt.Printf("Name: %s, Age: %d, Weight: %.2f\n", b.Name, b.Age, b.Weight)
}

type Dog struct {
	Animal // Dog 继承 Animal
}

func (d *Dog) Speak() {
	fmt.Println("Dog barks")
}

func (b *Dog) ShowInfo() {
	fmt.Printf("~~~~~Name: %s, Age: %d, Weight: %.2f\n", b.Name, b.Age, b.Weight)
}

func main() {
	var a Animal
	a.Speak()

	var d Dog = Dog{}
	y := &Dog{} // Dog{} 和 &Dog{} 区别
	d.Animal.Age = 10
	d.Animal.Name = "Buddy"
	d.Animal.Weight = 20.5
	d.Animal.ShowInfo() // 调用 Animal 的方法
	d.Speak()           // 调用 Dog 的 Speak 方法
	y.Animal.Age = 10
	y.Animal.Name = "Buddy"
	y.Animal.Weight = 20.5
	y.Animal.ShowInfo()
	y.Speak()

	//可以省略Animal
	d.Age = 5
	d.Name = "Max"
	d.Weight = 15.0
	d.ShowInfo() // 调用 Dog 的 ShowInfo 方法，这个取子
	//如果父 和子 都有相同的 结构，取就近原则
	d.Animal.ShowInfo()//这个取父
}
