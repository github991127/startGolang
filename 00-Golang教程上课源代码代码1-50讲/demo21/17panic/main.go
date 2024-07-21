package main

import (
	"fmt"
	"time"
)

//函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 50)
		fmt.Println("hello,world")
	}
}

//函数
func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	myMap[0] = "golang" //error
}

func main() {

	go sayHello()
	go test()

	//防止主进程退出这里使用time.Sleep演示，搭建也可以用sync.WaitGroup
	time.Sleep(time.Second)
}
