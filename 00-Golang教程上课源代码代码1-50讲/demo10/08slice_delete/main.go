package main

import "fmt"

func main() {

	// Go 语言中并没有删除切片元素的专用方法，我们可以使用切片本身的特性来删除元素
	a := []int{30, 31, 32, 33, 34, 35, 36, 37}
	// 要删除索引为 2 的元素  删除的元素：32  注意：append合并切片的时候最后一个元素要加...
	a = append(a[:2], a[3:]...)
	fmt.Println(a)

	//要删除35, 36
	sliceB := []int{30, 31, 32, 33, 34, 35, 36, 37}
	sliceB = append(sliceB[:5], sliceB[7:]...)
	fmt.Println(sliceB)

}
