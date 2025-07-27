package main

import "fmt"

// 方法的使用
type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age) //拿第一个变量 p.Name 填入 %s 的位置,拿第二个变量 p.Age 填入 %d 的位置
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	p.Greet()

}
