/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"fmt"
	"time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔50毫秒秒输出 "你好golang"
// 在主线程中也每隔50毫输出"你好golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行
func test() {
	for i := 0; i < 10; i++ {
		fmt.Println("test() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	go test() //表示开启一个协程
	for i := 0; i < 10; i++ {
		fmt.Println("main() 你好golang")
		time.Sleep(time.Millisecond * 100)
	}
}
