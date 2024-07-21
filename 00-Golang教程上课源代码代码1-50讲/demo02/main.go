/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

func main() {
	// fmt.Println("你好 golang")
	// fmt.Print("你好 golang")
	// fmt.Printf("你好 golang")

	/*
		fmt.Println("你好 golang")
		fmt.Print("你好 golang")
	*/

	//     ctrl+/ 单行注释       ctrl+/ 也可以解除单行注释

	/*
		Print Println区别
			1、一次输入多个值的时候 Println 中间有空格 Print 没有
			2、Println 会自动换行，Print 不会
	*/

	// fmt.Println("你好 golang")

	// fmt.Print("你好 golang")

	// fmt.Print("A")
	// fmt.Print("B")
	// fmt.Print("C")

	// fmt.Println("A")
	// fmt.Println("B")
	// fmt.Println("C")

	// fmt.Print("A", "B", "C")
	// fmt.Println("A", "B", "C")

	/*
		Println 和 Printf 区别
			1、Println默认会换行，Printf默认不会
			2、Printf是格式化输出 比Println功能更强大更灵活
	*/

	// var a = "aaa" //go语言中变量定义以后必须要使用

	// fmt.Println(a)

	// var a = "aaa" //go语言中变量定义以后必须要使用
	// fmt.Printf("%v", a)

	// var a int = 10
	// var b int = 3
	// var c int = 5
	// fmt.Println("a=", a, "b=", b, "c=", c)
	// fmt.Println("a=", a, "b=", b, "c=", c)

	// fmt.Printf("a=%v b=%v c=%v\n", a, b, c)    //   \n表示换行
	// fmt.Printf("a=%v b=%v c=%v", a, b, c)

	//类型推导方式定义变量
	a := 10
	b := 20
	c := 30
	fmt.Printf("a=%v b=%v c=%v\n", a, b, c)
	//使用Printf打印一个变量的类型
	fmt.Printf("a=%v a的类型是%T", a, a)

}
