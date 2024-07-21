package main

import "fmt"

// 闭包，有权访问另一个函数作用域中的变量的函数
func adder1() func() int {
	var i = 10
	return func() int {
		return i + 1
	} //闭包：函数里面嵌套一个函数 最后返回里面的函数
}

func adder2() func(int) int {
	var i = 10 //常驻内存，不污染全局
	return func(y int) int {
		i += y
		return i
	}
}

func main() {
	var fn = adder1()
	fmt.Println(fn()) //11
	fmt.Println(fn()) //11

	var fn2 = adder2()
	fmt.Println(fn2(10)) //20
	fmt.Println(fn2(10)) //30

}
