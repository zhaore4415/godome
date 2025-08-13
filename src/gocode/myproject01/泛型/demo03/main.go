package main

import "fmt"

//// 这两行是等价的（语义上）：
// func Print[T any](item T)
// func Print[T interface{}](item T)

type Stringable interface {
	String() string
}

func PrintString[T Stringable](v T) {
	fmt.Println(v.String())
}

// 定义一个结构体
type Person struct {
	Name string
	Age  int
}

// 实现 String() 方法
func (p Person) String() string {
	return fmt.Sprintf("Person: %s (age %d)", p.Name, p.Age)
}

// 另一个类型
type Book struct {
	Title string
	Year  int
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %q (%d)", b.Title, b.Year)
}

func main() {
	// ✅ 调用：传入 Person 实例
	p := Person{Name: "Alice", Age: 30}
	PrintString(p) // 输出: Person: Alice (age 30)

	// ✅ 调用：传入 Book 实例
	b := Book{Title: "Go Programming", Year: 2023}
	PrintString(b) // 输出: Book: "Go Programming" (2023)

}
