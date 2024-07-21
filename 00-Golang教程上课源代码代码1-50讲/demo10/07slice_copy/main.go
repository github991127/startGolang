package main

import "fmt"

func main() {
	/*
		值类型 ： 改变变量副本值的时候，不会改变变量本身的值
		引用类型：改变变量副本值的时候，会改变变量本身的值
	*/

	// 切片就是引用数据类型

	// sliceA := []int{1, 2, 3, 45}
	// sliceB := sliceA
	// sliceB[0] = 11
	// fmt.Println(sliceA)
	// fmt.Println(sliceB)

	/*
		[11 2 3 45]
		[11 2 3 45]
	*/

	//1、copy()函数复制切片

	sliceA := []int{1, 2, 3, 45}
	sliceB := make([]int, 4, 4)
	copy(sliceB, sliceA)

	fmt.Println(sliceA)
	fmt.Println(sliceB)
	sliceB[0] = 111
	fmt.Println(sliceA) //[1 2 3 45]
	fmt.Println(sliceB) //[111 2 3 45]

}
