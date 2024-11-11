package main

import "fmt"

// 先return后defer
// defer函数的执行顺序是先进后出的栈，即先注册的defer函数最后执行
func d1() int {
	x := 5
	defer func() {
		x++
	}()
	return x //5，defer函数最后执行，此之前的x值已经返回
}

func d2() (x int) {
	x = 5
	defer func() {
		x++
	}()
	//return //6,由于匿名返回非原子操作，返回值也会被defer函数修改
	return x //6,同上
}

func d3() (x int) {
	x = 5
	defer func(x int) {
		// x=0
		// fmt.Println(x)
		x++
	}(x) //defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	// fmt.Println("a=", x)
	return //5
}

// panic/recover进行异常处理，使用defer函数进行延迟调用
func errDiv(a int, b int) {
	defer func() {
		//panic("抛出一个异常") //手动抛出一个异常
		err := recover() //捕获自动或手动的panic异常
		if err != nil {
			fmt.Println("error:", err)
		}
	}()
	fmt.Println(a / b)
}
func main() {
	fmt.Println(d1())
	fmt.Println(d2())
	fmt.Println(d3())
	errDiv(10, 0) //error: runtime error: integer divide by zero
	errDiv(10, 4) //output: 2
}
