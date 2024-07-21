//decimal的用法：http://bbs.itying.com/topic/5e92a8938da5b70118ded75d
package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {

	// //加
	var num1 float64 = 3.1
	var num2 float64 = 4.2
	d1 := decimal.NewFromFloat(num1).Add(decimal.NewFromFloat(num2))
	fmt.Println(d1)

	//减去
	m1 := 8.2
	m2 := 3.8
	m3 := decimal.NewFromFloat(m1).Sub(decimal.NewFromFloat(m2))
	fmt.Println(m3)
	// 减法 Sub，乘法 Mul， 除法 Div 用法均与上述类似

}
