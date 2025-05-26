package main

import "fmt"

// 指针变量
func f5() {
	// golang 只有值传递，没有引用传递
	// golang 指针拷贝（类似引用传递效果）的有：slice,map,channel等等。其余为值拷贝
	var a = 10
	var p = &a //p指针变量   p的类型 *int（指针类型）
	fmt.Printf("a的值：%v;类型%T;地址%p\n", a, a, &a)
	fmt.Printf("p的值：%v;类型%T;地址%p\n", p, p, &p)
	fmt.Printf("*p的值：%v;类型%T;地址%p\n", *p, *p, &*p) //通过指针p访问a的值,*p就是a,指针是引用数据类型
}

// new 函数
func f6() {
	//实际开发中 new 函数不太常用，使用 new 函数得到的是一个指针类型，并且该指针对应的值为该类型的零值
	var a = new(int) //a是一个指针变量 类型是 *int的指针类型 指针变量对应的值是0
	fmt.Printf("a的值：%v 类型:%T 指针变量对应的值：%v", a, a, *a)
}

func main() {
	f5()
	f6()
}
