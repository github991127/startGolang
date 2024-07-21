/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

func main() {
	//在计算机底层 a 这个变量其实对应了一个内存地址
	// var a int = 10
	// fmt.Printf("a的值：%v a的类型%T a的地址%p", a, a, &a) //a的值：10 a的类型int a的地址0xc0000100a8

	// 指针也是一个变量，但它是一种特殊的变量，它存储的数据不是一个普通的值，而是另 一个变量的内存地址
	// var a = 10
	// var p = &a //p指针变量   p的类型 *int（指针类型）
	// fmt.Printf("a的值：%v a的类型%T a的地址%p\n", a, a, &a)
	// fmt.Printf("p的值：%v p的类型%T", p, p)

	//每一个变量都有自己的内存地址
	var a = 10
	var p = &a //p指针变量   p的类型 *int（指针类型）
	fmt.Printf("a的值：%v a的类型%T a的地址%p\n", a, a, &a)
	fmt.Printf("p的值：%v p的类型%T p的地址%p", p, p, &p)

}
