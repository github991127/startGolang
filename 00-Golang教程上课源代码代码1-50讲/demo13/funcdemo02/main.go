package main

import "fmt"

func sumFn(x, y int) int {
	return x + y
}

//return 关键词一次可以返回多个值
func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub
}

//返回值命名: 函数定义时可以给返回值命名，并在函数体中直接使用这些变量，最后通过 return 关键字 返回。

func calc1(x, y int) (sum int, sub int) {
	fmt.Println(sum, sub)
	sum = x + y
	sub = x - y
	fmt.Println(sum, sub)
	return
}

func calc2(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return
}

func test() {
	fmt.Println("执行...")
}

func main() {
	// sum1 := sumFn(10, 2)
	// fmt.Println(sum1)

	// a, b := calc(10, 2)
	// fmt.Println(a, b)

	// a, b := calc1(10, 2)
	// fmt.Println(a, "---", b)

	// a, b := calc2(10, 2)
	// fmt.Println(a, "---", b)

	// a, _ := calc2(10, 2)
	// fmt.Println(a)

	_, b := calc2(10, 2)
	fmt.Println(b)

	test()
}
