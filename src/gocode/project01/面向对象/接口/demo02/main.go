package main

import "fmt"

//接口嵌套
type SayHelloA interface {
	sayHelloA() //这里如果小写，下面的都要小写
}

type SayHelloB interface {
	sayHelloB() //这里如果小写，下面的都要小写
}

type SayHelloC interface {
	SayHelloA
	SayHelloB
	sayHelloC()
}
type Cat struct {
	name string
}

func (c Cat) sayHelloA() {
	fmt.Println("Hello from Cat A")
}

func (c Cat) sayHelloB() {
	fmt.Println("Hello from Cat B")
}
func (c Cat) sayHelloC() {
	fmt.Println("Hello from Cat C")
}

func main() {
	cat := Cat{name: "咪咪"}

	var c SayHelloC = cat //这个是包含了A和B接口的C接口

	c.sayHelloC()
	c.sayHelloA()
	c.sayHelloB()
	fmt.Println("Cat name:", cat.name)

}
