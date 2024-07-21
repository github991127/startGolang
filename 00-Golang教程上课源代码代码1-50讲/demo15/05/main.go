/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

func main() {

	// 实际想开发中 new 函数不太常用，使用 new 函数得到的是一个指类型针，并且该指针对应的值为该类型的零值
	// var a = new(int) //a是一个指针变量 类型是 *int的指针类型 指针变量对应的值是0

	// fmt.Printf("值：%v 类型:%T 指针变量对应的值：%v", a, a, *a) //值：0xc0000a0090 类型:*int 指针变量对应的值：0

	/*
		错误的写法
		var a *int //指针也是引用数据类型
		*a = 100
		fmt.Println(*a)
	*/

	//new方法给指针变量分配存储空间
	// var b *int
	// b = new(int)
	// *b = 100
	// fmt.Println(*b) //	fmt.Println(*a)

	var f = new(bool)
	fmt.Println(*f) //false
}
