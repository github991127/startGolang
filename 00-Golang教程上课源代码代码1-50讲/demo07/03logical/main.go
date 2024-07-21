//逻辑运算
package main

import "fmt"

//定义一个方法
func test() bool {
	fmt.Println("test...")
	return true
}
func main() {

	/*
		&&	逻辑 AND 运算符。 如果两边的操作数都是 True，则为 True，否则为 False。
		||	逻辑 OR 运算符。 如果两边的操作数有一个 True，则为 True，否则为 False。
		!	逻辑 NOT 运算符。 如果条件为 True，则为 False，否则为 True。
	*/

	// var a = 23
	// var b = 8
	// fmt.Println(a > 10 && b < 10) //true
	// fmt.Println(a > 24 && b < 10) //false
	// fmt.Println(a > 5 && b < 6)   //false
	// fmt.Println(a == 5 && b < 6) //false

	// var a = 23
	// var b = 8
	// fmt.Println(a > 10 || b < 10) //true
	// fmt.Println(a > 24 || b < 10) //true
	// fmt.Println(a > 5 || b < 6)   //true
	// fmt.Println(a == 5 || b < 6)  //false

	// flag := false
	// fmt.Println(!flag)

	//逻辑与和逻辑或短路
	//逻辑与：前面是false后面就不会执行  逻辑或：前面是true后面就不会执行

	/*
		test...
		执行
	*/
	// var a = 10
	// if a > 9 && test() {
	// 	fmt.Println("执行")
	// }

	/*
		输出：
		空
	*/
	// var a = 10
	// if a > 11 && test() {
	// 	fmt.Println("执行")
	// }

	/*
		输出：
		test...
		执行
	*/
	// var a = 10
	// if a > 11 || test() {
	// 	fmt.Println("执行")
	// }

	/*
		输出：执行
	*/
	var a = 10
	if a < 11 || test() {
		fmt.Println("执行")
	}

}
