package main

import "fmt"

func main() {
	/*
		值类型 ：改变变量副本值的时候，不会改变变量本身的值
		引用类型：改变变量副本值的时候，会改变变量本身的值
	*/

	//值类型:基本数据类型 和 数组都是值类型
	// var a = 10
	// b := a
	// a = 20
	// fmt.Println(a, b)

	// var arr1 = [...]int{1, 2, 3}
	// arr2 := arr1
	// arr1[0] = 11
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	//引用类型: 切片

	// var arr1 = []int{1, 2, 3}
	// arr2 := arr1
	// arr1[0] = 11
	// fmt.Println(arr1)
	// fmt.Println(arr2)

	var arr1 = []int{1, 2, 3, 4}
	arr2 := arr1
	arr2[0] = 111
	fmt.Println(arr1)
	fmt.Println(arr2)

}
