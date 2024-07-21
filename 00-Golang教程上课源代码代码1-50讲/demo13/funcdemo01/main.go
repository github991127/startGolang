/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

//求两个数的和
func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

//求两个数的差
func subFn(x int, y int) int {
	fmt.Println(x, y) //20 2
	sub := x - y
	return sub
}

//函数参数的简写
func subFn1(x, y int) int {
	sub := x - y
	return sub
}

//函数的可变参数，可变参数是指函数的参数数量不固定。Go 语言中的可变参数通过在参数名后加...来标识

func sumFn1(x ...int) int {

	// fmt.Printf("%v--%T", x, x) //[12 34 45 46]--[]int

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func sumFn2(x int, y ...int) int {
	fmt.Println(x, y) //100 [1 2 3 4]
	sum := x
	for _, v := range y {
		sum += v
	}
	return sum
}

func main() {
	// sum1 := sumFn(12, 3)
	// fmt.Println(sum1) //15

	// sum2 := sumFn(15, 5)
	// fmt.Println(sum2) //20

	// a := 20
	// b := 2
	// sub1 := subFn(a, b)
	// fmt.Println(sub1)

	// sub2 := subFn1(22, 4)
	// fmt.Println(sub2) //18

	// sum3 := sumFn1(12, 34, 45, 46) //137
	// fmt.Println(sum3)

	// sum4 := sumFn1(1, 2, 3, 4, 5) //15
	// fmt.Println(sum4)

	sum5 := sumFn2(100, 1, 2, 3, 4)
	fmt.Println(sum5)

}
