package main

import "fmt"

// 指针变量
func f5() {
	var a = 10
	var p = &a //p指针变量   p的类型 *int（指针类型）
	fmt.Printf("a的值：%v a的类型%T a的地址%p\n", a, a, &a)
	fmt.Printf("p的值：%v p的类型%T p的地址%p\n", p, p, &p)
	fmt.Printf("*p的值：%v *p的类型%T *p的地址%p\n", *p, *p, &*p) //通过指针p访问a的值,*p就是a,指针是引用数据类型

}

// new 函数
func f6() {
	//实际想开发中 new 函数不太常用，使用 new 函数得到的是一个指类型针，并且该指针对应的值为该类型的零值
	var a = new(int)                               //a是一个指针变量 类型是 *int的指针类型 指针变量对应的值是0
	fmt.Printf("值：%v 类型:%T 指针变量对应的值：%v", a, a, *a) //值：0xc0000a0090 类型:*int 指针变量对应的值：0

}
func main() {
	f5()
	f6()
}
