/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

//自定义类型
type myInt int

// type myFn func(int, int) int

//类型别名
type myFloat = float64

func main() {

	var a myInt = 10

	fmt.Printf("%v %T\n", a, a) //10 main.myInt

	var b myFloat = 12.3
	fmt.Printf("%v %T", b, b) //12.3 float64
}
