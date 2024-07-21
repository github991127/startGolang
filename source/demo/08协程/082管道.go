package main

import "fmt"

// 管道
func f3() {
	ch1 := make(chan int, 10)
	//ch2 := make(chan<- int, 10) //只写,常用于函数参数引用传递规范
	//ch3 := make(<-chan int, 10) //只读,常用于函数参数引用传递规范
	ch1 <- 2
	ch1 <- 3
	ch1 <- 5
	x := <-ch1
	fmt.Println(x)
	fmt.Printf("值：%v 容量：%v 长度%v\n", ch1, cap(ch1), len(ch1))
	ch2 := ch1
	fmt.Printf("值：%v 容量：%v 长度%v\n", ch2, cap(ch2), len(ch2)) // 管道是引用传递

}

// 管道遍历
func f4() {
	var ch1 = make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch1 <- i
	}
	close(ch1) //关闭管道
	//for range遍历，通道被关闭时会退出for range，如果没有关闭管道就会报个错误
	for v := range ch1 { //管道没有key
		fmt.Println(v)
	}

}
func main() {
	//f3() // 管道
	f4() // 管道遍历

}
