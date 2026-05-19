package main

import (
	"fmt"
)

func main() {
	fmt.Println("main start")
	a()
	fmt.Println("main end")
}

func a() {
	defer fmt.Println("defer in a")
	b()
	fmt.Println("a after b")
}

func b() {
	defer fmt.Println("defer in b")
	return
	fmt.Println("b after panic")
}
