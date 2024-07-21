package main

import (
	"fmt"
	"time"
)

func main() {
	// 在某些场景下我们需要同时从多个通道接收数据,这个时候就可以用到golang中给我们提供的select多路复用

	//1.定义一个管道 10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}
	//2.定义一个管道 5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}
	//使用select来获取channel里面的数据的时候不需要关闭channel
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
			return //注意退出...
		}
	}

}
