//位运算符
package main

import "fmt"

func main() {
	/*
		& 两位均为1才为1
		| 两位有一个为1就为1
		^ 相异或  两位不一样则为1
		<<  左移n位就是乘以2的n次方。  “a<<b”是把a的各二进位全部左移b位，高位丢弃，低位补0。
		>> 右移n位就是除以2的n次方。
	*/

	var a = 5 //  101
	var b = 2 //  010

	fmt.Println("a&b=", a&b)   // 000   值0
	fmt.Println("a|b=", a|b)   // 111  值7
	fmt.Println("a^b=", a^b)   // 111  值7
	fmt.Println("a<<b=", a<<b) // 10100  值20   5*2的2次方
	fmt.Println("a>>b=", a>>b) // 值1    5/2的2次方

}
