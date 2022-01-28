package conf

import (
	"fmt"
	"testing"
	"time"
)

func TestXxx(t *testing.T) {
	fmt.Println("test")
	// 构建一个通道
	ch := make(chan int)

	go func() {
		fmt.Println("go func1")
		ch <- 10
	}()

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("go func2")
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i
			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()

	select {
	case ch <- 1:
		fmt.Println("---")
	}
}
