package main

import "fmt"

//结构体 String()的使用
type person struct {
	name string
	age  int
}

func (s person) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", s.name, s.age)
}

func main() {
	p := person{
		name: "Alice",
		age:  30,
	}
	fmt.Println(p.String())
}
