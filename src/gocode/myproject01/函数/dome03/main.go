//可变参数

package main

import "fmt"

func calc(num1 ...int) {
	for _, v := range num1 {
		fmt.Println(v)
	}
}

func main() {

	var a int = 10
	calc(a, 20, 30, 40)

}
