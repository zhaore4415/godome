package main

import "fmt"

//结构体
type Person struct {
	Name string
	Age  int
}

func main() {
	var t *Person = new(Person)             // 声明一个Person类型的指针变量t，并使用new函数分配内存
	fmt.Println(t)                          // 输出为&{ 0}，因为指针指向的	结构体的字段都被初始化为零值
	t.Name = "Alice"                        // 修改指针指向的结构体的字段值
	t.Age = 30                              // 修改指针指向的结构体的字段值
	fmt.Println("Name:", t.Name)            // 输出 Name: Alice
	fmt.Println("Age:", t.Age)              // 输出 Age: 30
	(*t).Name = "Bob"                       // 修改指针指向的结构体的字段值  (*t).Name 与 t.Name 等价
	(*t).Age = 35                           // 修改指针指向的结构体的字段值
	fmt.Println("Updated Name:", (*t).Name) // 输出 Updated Name: Bob
	fmt.Println("Updated Age:", (*t).Age)   // 输出 Updated Age: 35

}
