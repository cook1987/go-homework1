package main

import (
	"fmt"
	"sync"
	"time"
)

// 题目 1
func testchannel() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 0)
	go func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			fmt.Println("通道接收的数字：", i)
		}
	}(ch)

	go func(ch chan<- int) {
		defer wg.Done()
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch)

	// 等待所有工作者完成
	wg.Wait()
}

// 题目二
func testBufferChannel() {
	var wg sync.WaitGroup
	wg.Add(2)
	ch := make(chan int, 10)
	// 生产者
	go func(ch chan<- int) {
		defer wg.Done()
		for i := 1; i <= 100; i++ {
			fmt.Println("生产者写入之前：", i)
			ch <- i
			fmt.Println("生产者写入之后：", i)
		}
		close(ch)
	}(ch)
	// 消费者
	go func(ch <-chan int) {
		defer wg.Done()
		for i := range ch {
			fmt.Println("消费者读取：", i)
			time.Sleep(50 * time.Millisecond)
		}
	}(ch)
	// 等待所有工作者完成
	wg.Wait()
}

// func main() {
// 	fmt.Println("----------------题目 1-------------")
// 	testchannel()
// 	fmt.Println("----------------题目 2-------------")
// 	testBufferChannel()
// }
