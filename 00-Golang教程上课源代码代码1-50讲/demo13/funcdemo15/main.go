package main

import "fmt"

/*
panic/recover

Go 语言中目前（Go1.12）是没有异常机制，但是使用 panic/recover 模式来处理错误。

panic 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效
*/

func fn1() {
	fmt.Println("fn1")
}
func fn2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("err:", err)
		}
	}()
	panic("抛出一个异常")
}
func main() {
	fn1()
	fn2()
	fmt.Println("结束")
}
