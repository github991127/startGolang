package main

import (
	"fmt"
	"strconv"
)

func main() {

	//1、通fmt.Sprintf() 把其他类型转换成 String 类型

	//注意：Sprintf 使用中需要注意转换的格式 int 为%d  float 为%f  bool 为%t   byte 为%c

	/*
		var i int = 20
		var f float64 = 12.456
		var t bool = true
		var b byte = 'a'

		str1 := fmt.Sprintf("%d", i)
		fmt.Printf("值：%v 类型：%T\n", str1, str1)

		str2 := fmt.Sprintf("%.2f", f)
		fmt.Printf("值：%v 类型：%T\n", str2, str2)

		str3 := fmt.Sprintf("%t", t)
		fmt.Printf("值：%v 类型：%T\n", str3, str3)

		str4 := fmt.Sprintf("%c", b)
		fmt.Printf("值：%v 类型：%T\n", str4, str4)
	*/

	//2、通过strconv  把其他类型转换成string类型

	/*
		FormatInt
		参数1：int64 的数值
		参数2：传值int类型的进制
	*/
	// var i int = 20
	// str1 := strconv.FormatInt(int64(i), 10)
	// fmt.Printf("值：%v 类型：%T\n", str1, str1)

	/*
		参数 1：要转换的值
		参数 2：格式化类型 'f'（-ddd.dddd）、
			 'b'（-ddddp±ddd，指数为二进制）、
			 'e'（-d.dddde±dd，十进制指数）、
			 'E'（-d.ddddE±dd，十进制指数）、
			 'g'（指数很大时用'e'格式，否则'f'格式）、
			 'G'（指数很大时用'E'格式，否则'f'格式）。
		 参数 3: 保留的小数点 -1（不对小数点格式化）
		 参数 4：格式化的类型 传入 64  32
	*/
	// var f float32 = 20.231313
	// str2 := strconv.FormatFloat(float64(f), 'f', 4, 32)
	// fmt.Printf("值：%v 类型：%T\n", str2, str2)

	// str3 := strconv.FormatBool(true) //没有任何意义
	// fmt.Printf("值：%v 类型：%T\n", str3, str3)

	a := 'b' //没有任何意义
	str4 := strconv.FormatUint(uint64(a), 10)
	fmt.Printf("值：%v 类型：%T\n", str4, str4) //值：98 类型：string
}
