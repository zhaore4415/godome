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

func (a *Cat) changGe() {
	fmt.Println("Cat changGe 唱歌")
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

	//断言 只有猫里面有这个方法
	//如果是猫，就调用它的changGe方法
	if cat, ok := s.(*Cat); ok {
		cat.changGe()
	}
	// 另一种方式
	// 断言类型
	// cat, ok := s.(*Cat)
	// if ok {
	// 	cat.changGe()
	// }
}

func main() {
	cat := &Cat{Name: "咪咪", Age: 5, Weight: 10.2}

	MakeSound(cat)

}
