package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 题目 1
type shareCal struct {
	mu  sync.Mutex
	num int
}

func (s *shareCal) incrment() {
	s.mu.Lock()
	s.num++
	s.mu.Unlock()
}

func cal() {
	var wg sync.WaitGroup
	wg.Add(10)
	share := shareCal{}
	for i := 0; i < 10; i++ {
		go func(s *shareCal) {
			for i := 0; i < 1000; i++ {
				s.incrment()
			}
			wg.Done()
		}(&share)
	}
	// 等待所有工作者完成
	wg.Wait()
	fmt.Println("计算结果：", share.num)
}

// 题目 2
type shareCal2 struct {
	num atomic.Int64
}

func (s *shareCal2) incrment() {
	s.num.Add(1)
}

func cal2() {
	var wg sync.WaitGroup
	wg.Add(10)
	share := shareCal2{}
	for i := 0; i < 10; i++ {
		go func(s *shareCal2) {
			for i := 0; i < 1000; i++ {
				s.incrment()
			}
			wg.Done()
		}(&share)
	}
	// 等待所有工作者完成
	wg.Wait()
	fmt.Println("计算结果：", share.num.Load())
}

func main() {
	fmt.Println("----------------题目 1-------------")
	cal()
	fmt.Println("----------------题目 2-------------")
	cal2()
}
