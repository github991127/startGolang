/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

func main() {
	//1、定义float类型

	// var a float32 = 3.12
	// fmt.Printf("值：%v--%f,类型%T\n", a, a, a) //3.12--3.120000,类型float32
	// fmt.Println(unsafe.Sizeof(a))          //float32占用4个字节

	// var b float64 = 3.12
	// fmt.Printf("值：%v--%f,类型%T\n", b, b, b) //值：3.12--3.120000,类型float64
	// fmt.Println(unsafe.Sizeof(b))          //float64占用8个字节

	//2、 %f  输出float类型   %.2f 输出数据的时候保留2位小数

	// var c float64 = 3.1415925535
	// fmt.Printf("%v--%f--%.2f", c, c, c)

	// 保留4位小数
	// var c float64 = 3.1415926535
	// fmt.Printf("%.4f", c)

	//3、 64位的系统中 浮点数默认是 float64

	// f1 := 3.1345656456

	// fmt.Printf("%f--%T", f1, f1)

	//4、Golang 科学计数法表示浮点类型

	// var f2 float32 = 3.14e2 //表示f2等于3.14*10的2次方
	// fmt.Printf("%v--%T", f2, f2)

	// var f3 float32 = 3.14e-2 //表示f3等于3.14除以10的2次方
	// fmt.Printf("%v--%T", f3, f3)

	// 5、Golang 中 float 精度丢失问题

	// var f4 float64 = 1129.6
	// fmt.Println(f4 * 100) //输出：112959.99999999999
	// m1 := 8.2
	// m2 := 3.8
	// fmt.Println(m1 - m2) //期望是 4.4，结果打印出了 4.399999999999999

	// 6、int类型转换成float类型

	// a := 10
	// b := float64(a)
	// fmt.Printf("a的类型是%T,b的类型是%T", a, b)

	// 7、int类型转换成float类型

	// var a1 float32 = 23.4
	// a2 := float64(a1)
	// fmt.Printf("a1的类型是%T,a2的类型是%T", a1, a2)

	//8、float类型转换成int类型 （不建议这样做）

	var c1 float32 = 23.45

	c2 := int(c1)

	fmt.Printf("c2的值：%v  c2的类型：%T", c2, c2)

}
