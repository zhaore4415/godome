
//使用互斥锁 同步协程
package main

import (
    "fmt"
    "sync"
    "time"
)

// 定义一个结构体来模拟共享资源
type SafeCounter struct {
    mu    sync.Mutex // 互斥锁
    value int        // 共享的计数值
}

// Increment 方法用于安全地增加计数器的值
func (sc *SafeCounter) Increment() {
    sc.mu.Lock() // 锁定
    defer sc.mu.Unlock() // 在函数退出时解锁
    sc.value++
    fmt.Printf("Incremented to %d\n", sc.value)
}

// Value 方法返回当前的计数值
func (sc *SafeCounter) Value() int {
    sc.mu.Lock()
    defer sc.mu.Unlock()
    return sc.value
}

func worker(id int, wg *sync.WaitGroup, sc *SafeCounter) {
    defer wg.Done()
    for i := 0; i < 3; i++ {
        time.Sleep(time.Duration(100*id) * time.Millisecond) // 模拟不同的工作时间
        sc.Increment()
    }
}

func main() {
    var wg sync.WaitGroup
    counter := SafeCounter{}

    for i := 1; i <= 5; i++ {
        wg.Add(1)
        go worker(i, &wg, &counter)
    }

    wg.Wait()

    fmt.Println("Final counter:", counter.Value())
}