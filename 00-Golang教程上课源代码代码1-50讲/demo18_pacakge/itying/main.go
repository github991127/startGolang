package main

import (
	"fmt"
	"itying/calc"
	"itying/tools"
)
func main() {
	sum := calc.Add(10, 2)
	fmt.Println(sum) //12
	// fmt.Println(calc.aaa)
	sub := calc.Sub(10, 2)
	fmt.Println(sub)

	//调用tools包里面的方法
	b := tools.Mul(2, 6)
	fmt.Println(b)

}
