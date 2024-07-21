//赋值运算
package main

import "fmt"

func main() {
	/*
		=	简单的赋值运算符，将一个表达式的值赋给一个左值
		+=	相加后再赋值
		-=	相减后再赋值
		*=	相乘后再赋值
		/=	相除后再赋值
		%=	求余后再赋值
	*/

	// var a = 20
	// a = 21
	// fmt.Println(a)

	// var a = 23 + 2
	// fmt.Println(a)

	// var a = 10
	// b := a
	// fmt.Println(b)

	// var a = 10
	// b := a + 2
	// fmt.Println(b)

	// var a = 10
	// a += 3 // 等价于a=a+3
	// fmt.Println(a)

	// var a = 10
	// a = a + 3
	// fmt.Println(a)

	// var a = 10
	// a -= 3 //a=a-3
	// fmt.Println(a)

	// var a = 10
	// a *= 3 //a=a*3
	// fmt.Println(a)

	// var a = 10
	// a /= 3         // a=a/3
	// fmt.Println(a) //结果：3

	// var a float64 = 10
	// a /= 3         // a=a/3
	// fmt.Println(a) //结果：3.3333333333333335

	var a = 10
	a %= 3         // a=a%3
	fmt.Println(a) //1
}
