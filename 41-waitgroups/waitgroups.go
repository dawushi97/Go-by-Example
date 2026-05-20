package main

import (
	"fmt"
	"sync"
	"time"
)

// TODO 1: 实现 worker(id int) 函数
// - 打印 "Worker X starting"
// - time.Sleep(time.Second) 模拟耗时任务
// - 打印 "Worker X done"
func worker(id int) {
	fmt.Printf("worker %d startting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func main() {
	// TODO 2: 声明一个 sync.WaitGroup
	// - var wg sync.WaitGroup
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Go(func() {
			worker(i)
		})
	}
	/* 	for i := 0; i < 5; i++ {
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker(i)
	}()
	} */
	wg.Wait()
}
