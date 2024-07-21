package main

import "fmt"

/*
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。
*/

func f1() {
	fmt.Println("开始")
	defer func() {
		fmt.Println("aaaa")
		fmt.Println("bbbb")
	}()

	fmt.Println("结束")
}

//	调用方法返回的是0   fmt.Println(f2()) //0
func f2() int {
	var a int //0
	defer func() {
		a++
	}()
	// fmt.Println("结束")
	return a
}

func f3() (a int) {
	defer func() {
		a++
	}()
	return a
}

func main() {

	//1、defer的使用演示
	// fmt.Println("开始")
	// defer fmt.Println(1)
	// defer fmt.Println(2)
	// defer fmt.Println(3)
	// fmt.Println("结束")

	//2、defer在命名返回值和匿名返回 函数中表现不一样

	// f1()

	fmt.Println(f2()) //0

	fmt.Println(f3()) //1

}
