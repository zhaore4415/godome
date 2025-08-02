package main

//协程与管道案例 ,加放写和读，加入WaitGroup	
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	intChan := make(chan int, 3)

	// 启动写协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			intChan <- i
			time.Sleep(time.Second)
		}
		close(intChan)
	}()

	// 启动读协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range intChan {
			fmt.Println("Received:", v)
		}
	}()

	wg.Wait()
}
// 示例：使用管道和协程