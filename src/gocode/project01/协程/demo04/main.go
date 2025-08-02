package main

import (
	"fmt"
	"sync"
	"time"
)

// 读写锁
var lock sync.RWMutex

func read() {
	lock.RLock() //读锁
	defer lock.RUnlock()
	// 读取共享数据

	fmt.Println("Reading data")
	time.Sleep(100 * time.Millisecond) // 模拟读取操作
	fmt.Println("Read complete")

	lock.RUnlock()
}

func write() {
	lock.Lock() //写锁
	defer lock.Unlock()
	// 写入共享数据

	fmt.Println("Writing data")
	time.Sleep(200 * time.Millisecond) // 模拟写入操作
	fmt.Println("Write complete")

	lock.Unlock()
}

func main() {
	var wg sync.WaitGroup

	// 启动多个读协程
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			read()
		}(i)
	}

	// 启动一个写协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		write()
	}()

	wg.Wait() // 等待所有协程完成
	fmt.Println("All operations completed.")
}
