//算术运算
package main

import "fmt"

func main() {

	/*
		+	相加
		-	相减
		*	相乘
		/	相除
		%	求余=被除数-（被除数/除数）*除数

		1、除法注意：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分
		2、取余注意  余数=被除数-（被除数/除数）*除数
		3、注意： ++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。
	*/

	// var a = 6
	// var b = 3
	// fmt.Println(a + b)
	// fmt.Println(a - b)
	// fmt.Println(a * b)
	// var c = a * b
	// fmt.Println(c)

	// 1、除法注意：如果运算的数都是整数，那么除后，去掉小数部分，保留整数部分

	// var a = 10
	// var b = 3
	// fmt.Println(a / b) //3

	// var a = 10.0
	// var b = 3.0
	// fmt.Println(a / b) //3.3333333333333335

	// 2、取余注意    余数=被除数-（被除数/除数）*除数
	// var a = 10
	// var b = 3
	// fmt.Println(a % b) //1

	// fmt.Println(-10 % 3) // - 10 - (-10/3)*3  =-1

	// fmt.Println(10 % -3) // 10-(10/-3)*-3=1

	//3、注意： ++（自增）和--（自减）在Go语言中是单独的语句，并不是运算符。

	//1、注意：在 golang 中，++ 和 -- 只能独立使用 错误写法如下：

	/*
		var i int = 8
		var a int
		a = i++ //错误，i++只能独立使用
		a = i-- //错误, i--只能独立使用
		fmt.Ptintln(a)
	*/

	//2、注意：在 golang 中没有前++ 错误写法如下：

	// var a = 12
	// ++a  //错误写法
	// fmt.Println(a)

	//3、正确的写法

	var a = 12
	a++
	a--
	fmt.Println(a)

}
