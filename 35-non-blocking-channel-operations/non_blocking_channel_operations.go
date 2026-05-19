package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")

	}
	msg := "hi"
	done := make(chan bool, 1)
	go func(done chan bool) {
		ms := <-message
		fmt.Println("received message", ms)
		done <- true
	}(done)
	time.Sleep(time.Second)
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}
	<-done
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}
