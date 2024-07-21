package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x interface{}) {
	// var num = 10 + x //(mismatched types int and interface {}
	// b, _ := x.(int)
	// var num = 10 + b
	// fmt.Println(num) //23

	//反射来实现这个功能
	// v := reflect.ValueOf(x)
	// fmt.Println(v) //13
	// var n = v + 12
	// fmt.Println(n) //mismatched types reflect.Value and int

	//反射获取变量的原始值
	v := reflect.ValueOf(x)
	var m = v.Int() + 12
	fmt.Println(m) //25

}
func main() {

	var a = 13
	reflectValue(a)
}
