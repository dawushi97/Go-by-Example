package main

import (
	"errors"
	"fmt"
)

// - Go 惯例：error 永远是最后一个返回值，nil 表示无错误
func f(arg int) (int, error) {
		if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

// TODO 2: 定义两个 Sentinel Error（哨兵错误）
// - 命名惯例：以 Err 开头
// - 类比 C++：类似 errno 或预定义的 error_code
var ErrOutOfTea = errors.New("no more tea available")
var ErrPower = errors.New("can't boil water")

// TODO 3: 写一个函数 makeTea(arg int) error
//   - arg == 2 时返回 ErrOutOfTea
//   - arg == 4 时返回 fmt.Errorf("making tea: %w", ErrPower)
//     （%w 动词会"包装"错误，保留原始错误链）
//   - 其他情况返回 nil
func makeTea(arg int) error {
	if arg == 2 {
		return ErrOutOfTea
	} else if arg == 4 {
		return fmt.Errorf("making tea: %w", ErrPower)
	}
	return nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f(i); e != nil {
			fmt.Println("f failed:", e)
		} else {
			fmt.Println("f worked:", r)
		}
	}

	// - 用 errors.Is(err, ErrOutOfTea) 检查是否匹配哨兵错误
	// - 用 errors.Is(err, ErrPower) 检查（即使被 %w 包装过也能匹配！）
	// - 对比 C++：类似 catch 不同异常类型，但 Go 用返回值而非异常
	for i := range 5 {
		if err := makeTea(i); err != nil {
			if errors.Is(err, ErrOutOfTea) {
				fmt.Println("We should buy new tea!")
			} else if errors.Is(err, ErrPower) {
				fmt.Println("Now it is dark.")
			} else {
				fmt.Printf("unknown error: %s\n", err)
			}
			continue
		}
		fmt.Println("Tea is ready!")
	}
}
