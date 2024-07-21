package main

import "fmt"

//golang中空接口也可以直接当做类型来使用，可以表示任意类型

//空接口作为函数的参数
func show(a interface{}) {
	fmt.Printf("值:%v 类型:%T\n", a, a)
}
func main() {

	// var a interface{}
	// a = 20
	// fmt.Printf("值:%v 类型:%T\n", a, a)
	// a = "你好golang"
	// fmt.Printf("值:%v 类型:%T\n", a, a)
	// a = true
	// fmt.Printf("值:%v 类型:%T\n", a, a)

	// show(20)
	// show("你好golang")
	// slice := []int{1, 2, 34, 4}
	// show(slice)

	// var m1 = make(map[string]string)
	// m1["username"] = "张三"
	// m1["age"] = "20"
	// fmt.Println(m1)

	// var m1 = make(map[string]interface{})
	// m1["username"] = "张三"
	// m1["age"] = 20
	// m1["married"] = true
	// fmt.Println(m1)

	var s1 = []interface{}{1, 2, "你好", true} //[1 2 你好 true]
	fmt.Println(s1)
}
