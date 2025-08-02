package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // 在函数退出时，将 WaitGroup 的计数器减 1
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
	//wg.Done() 本来写在最后，等程序执行完，但是有了defer 这样就写在第一行，程序跑完才会执行他
}

func main() {
	var wg sync.WaitGroup
	numWorkers := 5

	// 启动多个协程
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1) // 将 WaitGroup 的计数器加 1
		go worker(i, &wg)
	}

	wg.Wait() // 阻塞，直到计数器变为 0
	fmt.Println("All workers completed.")
}
