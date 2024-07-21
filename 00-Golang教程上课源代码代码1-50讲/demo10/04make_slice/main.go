package main

import "fmt"

func main() {
	/*1、切片声明和初始化的几种方法：*/

	// var arr1 []int
	// fmt.Printf("%v - %T - 长度:%v\n", arr1, arr1, len(arr1))

	// var arr2 = []int{1, 2, 34, 45}
	// fmt.Printf("%v - %T - 长度:%v\n", arr2, arr2, len(arr2))

	// arr3 := []string{"php", "java"}
	// fmt.Printf("%v - %T - 长度:%v\n", arr3, arr3, len(arr3))

	// var arr4 = []int{1: 2, 2: 4, 5: 6}
	// fmt.Printf("%v - %T - 长度:%v", arr4, arr4, len(arr4))

	//2、make()函数创建一个切片  make([]T, size, cap)

	var sliceA = make([]int, 4, 8)
	// fmt.Println(sliceA) //[0 0 0 0]
	fmt.Printf("%d--%d", len(sliceA), cap(sliceA))

}
