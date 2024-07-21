package main

import "fmt"

//函数作为另一个函数参数
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

//自定义一个方法类型
type calcType func(int, int) int

func calc(x, y int, cb calcType) int {
	return cb(x, y)
}

func main() {
	// sum := calc(10, 5, add)
	// fmt.Println(sum) //15

	// s := calc(10, 5, sub)
	// fmt.Println(s) //5

	j := calc(3, 4, func(x, y int) int {
		return x * y
	})
	fmt.Println(j) //12

}
