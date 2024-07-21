/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

func main() {

	// var num1 float64 = 3.1
	// var num2 float64 = 4.2
	// fmt.Println(num1 + num2) //7.300000000000001

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

	f1 := 1.2
	f2 := 3.8

	f3 := decimal.NewFromFloat(f1).Mul(decimal.NewFromFloat(f2))
	fmt.Println(f3)

	const json = `{"name":{"first":"Janet","last":"四"},"age":47}`

	value := gjson.Get(json, "name.last")
	println(value.String())
}
