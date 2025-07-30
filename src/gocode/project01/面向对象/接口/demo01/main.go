package main

import "fmt"

type SayHello interface {
	sayHello() //这里如果小写，下面的都要小写
}

type Cat struct {
	Name   string
	Age    int
	Weight float64
}

func (a *Cat) sayHello() {
	fmt.Println("Cat speaks")
}

type Dog struct {
	Name   string
	Age    int
	Weight float64
}

func (a *Dog) sayHello() {
	fmt.Println("Dog speaks")
}

// 一个使用接口的函数
func MakeSound(s SayHello) {
	s.sayHello()
}

func main() {
	cat := &Cat{Name: "咪咪", Age: 5, Weight: 10.2}
	dog := &Dog{"小白", 6, 15.4}

	MakeSound(cat)
	MakeSound(dog)
}
