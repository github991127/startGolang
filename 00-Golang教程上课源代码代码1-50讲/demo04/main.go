/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

func main() {
	// var username = "张三"
	// username = "李四"
	// username = "王五"

	// fmt.Println(username)

	//1、go语言中的常量

	// const pi = 3.14159
	// fmt.Println(pi)

	// const A = "A"
	// A = "AAA"  //cannot assign to A  常量的值不可以改变
	// fmt.Println(A)

	//2、多个常量也可以一起声明

	// const (
	// 	A = "A"
	// 	B = "B"
	// )
	// fmt.Println(A, B)

	//3、const 同时声明多个常量时，如果省略了值则表示和上面一行的值相同。

	// const (
	// 	n1 = 100
	// 	n2
	// 	n3
	// 	n4
	// )
	// fmt.Println(n1, n2, n3, n4)

	/*
		4、iota 是 golang 语言的常量计数器,只能在常量的表达式中使用。


			iota 在 const 关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量 声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)
	*/

	// const a = iota
	// fmt.Println(a)

	// const (
	// 	n1 = iota
	// 	n2
	// 	n3
	// 	n4
	// )
	// fmt.Println(n1, n2, n3, n4) //0 1 2 3

	// const (
	// 	n1 = iota
	// 	_
	// 	n3
	// 	n4
	// )

	// fmt.Println(n1, n3, n4) //0  2 3

	//iota 声明中间插队
	// const (
	// 	n1 = iota //0
	// 	n2 = 100  //100
	// 	n3 = iota //2
	// 	n4        //3
	// )
	// fmt.Println(n1, n2, n3, n4)

	//多个 iota 定义在一行

	// const (
	// 	n1, n2 = iota + 1, iota + 2 //1  2
	// 	n3, n4                      //2 3
	// 	n5, n6                      //3 4
	// )

	// fmt.Println(n1, n2)
	// fmt.Println(n3, n4)
	// fmt.Println(n5, n6)

	//定义变量
	// n1, n2 := 20, 30
	// fmt.Println(n1, n2)

	//5、go语言中变量的名字是区分大小写的
	// var age = 20
	// var Age = 30
	// fmt.Println(age, Age)

	// var username = "zhangsan"

	// var DNS = "192.112.23.2"

}
