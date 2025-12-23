package main

import (
	"fmt"
	"time"
)

// 题目 1
func testgo() {
	go func() {
		for i := 1; i <= 10; i++ {
			if i%2 == 1 {
				fmt.Println("打印从1到10的奇数: ", i)
			}
		}
	}()
	go func() {
		for i := 2; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println("打印从2到10的偶数: ", i)
			}
		}
	}()
}

// 题目 2
func testgoSchedule(funclist []func()) {
	for i, taskd := range funclist {
		if taskd != nil {
			go func() {
				start := time.Now()
				taskd()
				duration := time.Since(start)
				fmt.Println("任务：", i, "，耗时：", duration)
			}()
		}
	}
}

// func main() {
// 	testgo()

// 	funclist := []func(){}
// 	funclist = append(funclist, func() {
// 		for i := 0; i < 10; i++ {
// 			fmt.Println("打印1： ", i)
// 		}
// 	})
// 	funclist = append(funclist, func() {
// 		for i := 10; i < 20; i++ {
// 			fmt.Println("打印2： ", i)
// 		}
// 	})
// 	testgoSchedule(funclist)

// 	time.Sleep(2 * time.Second)
// }
