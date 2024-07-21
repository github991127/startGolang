package main

import "fmt"

/*
Go语言中的defer语句会将其后面跟随的语句进行延迟处理。在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行，也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行。

问，上面代码的输出结果是？

提示：defer注册要延迟执行的函数时该函数所有的参数都需要确定其值

*/

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

/*
注册顺序
	defer calc("AA", x, calc("A", x, y))
	defer calc("BB", x, calc("B", x, y))
执行顺序
	1、defer calc("BB", x, calc("B", x, y))
	2、defer calc("AA", x, calc("A", x, y))

*/

//1、calc("A", x, y)     A 1 2 3
//2、calc("B", x, y)     B 10 2 12
//3、calc("BB", x, calc("B", x, y))    BB 10 12  22
//4、calc("AA", x, calc("A", x, y))    AA 1 3  4
