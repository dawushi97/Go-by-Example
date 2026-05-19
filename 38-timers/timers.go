package main

import (
	"fmt"
	"time"
)

func main() {
	// TODO 1: 创建一个 2 秒的 Timer
	// - 使用 time.NewTimer(2 * time.Second)
	// - <-timer1.C 阻塞等待定时器触发（C 是一个 chan Time）
	// - 打印 "Timer 1 fired"
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("Timer1 fired")

	// TODO 2: 创建一个 1 秒的 Timer，然后在它触发前 Stop 它
	// - 用 go func() 在 goroutine 中等待 <-timer2.C
	// - 在主 goroutine 中调用 timer2.Stop()
	// - Stop() 返回 true 表示成功取消（定时器还没触发）
	// - 打印 "Timer 2 stopped"
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer2 stopped")
	}
	// TODO 3: Sleep 2 秒，确认 timer2 确实没有触发
	// - 提示：time.Sleep vs time.NewTimer 的区别：
	//   Timer 可以被取消(Stop)，Sleep 不能
	time.Sleep(2 * time.Second)

}
