package main

import "fmt"

func main() {
	//1、创建channel
	ch := make(chan int, 3)

	//2、给管道里面存储数据
	ch <- 10
	ch <- 21
	ch <- 32

	//3、获取管道里面的内容
	a := <-ch
	fmt.Println(a) //10

	<-ch //从管道里面取值   //21
	c := <-ch
	fmt.Println(c) //32
	ch <- 56
	ch <- 66
	//4、打印管道的长度和容量
	fmt.Printf("值：%v 容量：%v 长度%v\n", ch, cap(ch), len(ch)) //值：0xc0000d0080 容量：3 长度2

	// 5、管道的类型（引用数据类型）

	ch1 := make(chan int, 4)

	ch1 <- 34
	ch1 <- 54
	ch1 <- 64

	ch2 := ch1
	ch2 <- 69
	<-ch1
	<-ch1
	<-ch1
	d := <-ch1
	fmt.Println(d) //69

	//8、管道阻塞

	// ch6 := make(chan int, 1)
	// ch6 <- 34
	// ch6 <- 64 //all goroutines are asleep - deadlock!

	// 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报告 deadlock
	// ch7 := make(chan string, 2)
	// ch7 <- "数据1"
	// ch7 <- "数据2"
	// m1 := <-ch7
	// m2 := <-ch7
	// m3 := <-ch7
	// fmt.Println(m1, m2, m3) //fatal error: all goroutines are asleep - deadlock!

	//正确的写法
	ch8 := make(chan int, 1)
	ch8 <- 34
	<-ch8
	ch8 <- 67
	<-ch8
	ch8 <- 78
	m4 := <-ch8
	fmt.Println(m4)

}
