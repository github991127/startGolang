package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 主线程退出后所有的协程无论有没有执行完毕都会退出，所以我们在主进程中可以通过WaitGroup等待协程执行完毕
var wg sync.WaitGroup

func test1() {
	for i := 0; i < 10; i++ {
		fmt.Println("test1() 你好golang-", i)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done() //协程计数器-1
}

func test2() {
	for i := 0; i < 10; i++ {
		fmt.Println("test2() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}
	wg.Done() //协程计数器-1
}

// 主线程开启协程
func f1() {
	wg.Add(1)  //协程计数器+1
	go test1() //开启协程
	wg.Add(1)  //协程计数器+1
	go test2()

	for i := 0; i < 10; i++ {
		fmt.Println("test3() 你好golang-", i)
		time.Sleep(time.Millisecond * 100)
	}

	wg.Wait() //等待协程执行完毕
	fmt.Println("主线程退出...")
}

// 设置使用多个cpu
func f2() {
	//获取当前计算机上面的Cup个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以自己设置使用多个cpu
	runtime.GOMAXPROCS(cpuNum - 1)
}

func main() {
	//f1() // 主线程开启协程
	f2() // 设置使用多个cpu
}
