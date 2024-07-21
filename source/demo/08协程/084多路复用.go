package main

import (
	"fmt"
	"time"
)

// 多路复用
func f5() {
	// select,同时从多个通道接收数据

	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	
	//select,不需要关闭channel
	for {
		select {
		case v := <-intChan:
			fmt.Printf("从 intChan 读取的数据%d\n", v)
			time.Sleep(time.Millisecond * 50)
		case v := <-stringChan:
			fmt.Printf("从 stringChan 读取的数据%v\n", v)
			time.Sleep(time.Millisecond * 50)
		default:
			fmt.Printf("数据获取完毕")
			return //注意退出循环
		}
	}
}

func main() {
	f5() //多路复用
}
