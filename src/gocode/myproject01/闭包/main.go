package main

import "fmt"

//实现闭包
func getSum() func(int) int {
	var sum int = 0
	return func(num int) int {
		sum = sum + num
		return sum
	}

}

func main() {

	ff := getSum()
	fmt.Println(ff(1))
	fmt.Println(ff(2))
	fmt.Println(ff(2)) // 这种可以一直累加

	fmt.Println("------------------")
	fmt.Println(getSum2(1)) // 这里每次从0开始
	fmt.Println(getSum2(2))
}

func getSum2(num int) int {
	var sum int = 0
	sum = sum + num
	return sum
}
