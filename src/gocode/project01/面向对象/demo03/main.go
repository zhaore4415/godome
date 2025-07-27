package main

import "fmt"

// 结构体之间的转换
type Person struct {
	Name string
	Age  int
}

type Employee struct {
	FullName string
	Age      int
	Position string
}

type Employee2 struct {
	FullName string
	Age      int
	Position string
}

//结构体取个别名
type Emp = Employee
type Emp2 = Employee2

func main() {
	p := Person{Name: "Alice", Age: 30}
	e := Employee{
		FullName: p.Name,
		Age:      p.Age,
		Position: "Developer",
	}
	fmt.Println("Employee:", e)

	// 如果是相同的字段名，可以使用类型转换来简化结构体之间的转换。例如：
	var a Employee2 = Employee2{"5", 30, "Developer"}
	var b Employee = Employee{"5", 30, "Manager"}
	a = Employee2(b) // 直接转换，假设字段名和类型都匹配
	fmt.Println("Employee2:", a)
	fmt.Println("Employee2-1:", b)

	// 使用别名
	var emp1 Emp = Employee{"Alice", 30, "Developer"}
	var emp2 Emp2 = Employee2{"Bob", 25, "Designer"}
	fmt.Println("Employee1:", emp1)
	fmt.Println("Employee2:", emp2)

}
