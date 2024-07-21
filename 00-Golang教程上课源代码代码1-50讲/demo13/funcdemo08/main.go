package main

import "fmt"

//函数作为返回值
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

//定义一个方法类型
type calcType func(int, int) int

func do(o string) calcType {
	switch o {
	case "+":
		return add
	case "-":
		return sub
	case "*":
		return func(x, y int) int {
			return x * y
		}
	default:
		return nil
	}
}

func main() {
	// var a = do("+")
	// fmt.Println(a(12, 4)) //16

	b := do("*")
	fmt.Println(b(3, 4)) //12
}
