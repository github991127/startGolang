/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

func main() {
	/*
		Go语言中以bool类型进行声明布尔型数据，布尔型数据只有true（真）和false（假）两个值。

		注意：
		1.布尔类型变量的默认值为false。
		2.Go 语言中不允许将整型强制转换为布尔型.
		3.布尔型无法参与数值运算，也无法与其他类型进行转换。
	*/

	// var flag = true

	// fmt.Printf("%v--%T", flag, flag)

	//1.布尔类型变量的默认值为false。
	// var b bool
	// fmt.Printf("%v", b)

	//2.string型变量的默认值为空。

	// var s string
	// fmt.Printf("%v", s)

	//3.int型变量的默认值为0。

	// var i int
	// fmt.Printf("%v", i)

	//4.float型变量的默认值为0。

	// var f float32
	// fmt.Printf("%v", f)

	//5、Go 语言中不允许将整型强制转换为布尔型.
	// var a = 1
	// if a { //错误写法
	// 	fmt.Printf("true")
	// }

	//6.布尔型无法参与数值运算，也无法与其他类型进行转换。

	// var s = "this is str"
	// if s { //错误写法
	// 	fmt.Printf("true")
	// }

	var f1 = false
	if f1 { //正确写法
		fmt.Printf("true")
	} else {
		fmt.Printf("false")
	}
}
