package main

import "fmt"

type calc func(int, int) int //表示定义一个calc的类型

type myInt int

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}

func test() {
	fmt.Println("test...")
}

func main() {

	// var c calc
	// c = sub
	// fmt.Printf("c的类型：%T\n", c) //c的类型：main.calc
	// fmt.Println(c(10, 5))      //5

	// f := sub
	// fmt.Printf("c的类型：%T\n", f) //c的类型：func(int, int) int
	// fmt.Println(f(15, 5))      //10

	var a int = 10
	var b myInt = 20
	fmt.Printf("a的类型：%T  b的类型：%T\n", a, b) //a的类型：int  b的类型：main.myInt
	fmt.Println(a + int(b))                //30
}
