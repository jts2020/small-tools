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
		ch <- 10
	}()

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("go func")
		// 从3循环到0
		for i := 3; i >= 0; i-- {
			// 发送3到0之间的数值
			ch <- i
			// 每次发送完时等待
			time.Sleep(time.Second)
		}
	}()
	fmt.Println("main func")
	// 遍历接收通道数据
	for data := range ch {
		// 打印通道数据
		fmt.Println(data)
		// 当遇到数据0时, 退出接收循环
		if data == 0 {
			break
		}
	}

}
