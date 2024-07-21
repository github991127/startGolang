package main

import "fmt"

// 空接口定义
type A interface{}

func show(a interface{}) {
	fmt.Printf("值:%v 类型:%T\n", a, a)

	// 判断类型,接口名.(类型)，返回：value和bool
	if _, ok := a.(int); ok {
		fmt.Println("int类型")
	} else {
		fmt.Println("其他类型")
	}

	// 求类型，接口名.(type)
	switch a.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("传入错误")
	}
}

func main() {
	// 应用1：空接口作为函数的参数，可以接收任意类型的值
	var a A
	var str = "你好golang"
	a = str //让字符串实现A这个接口
	show(a)

	// 应用2：空接口实现多类型的切片或map
	var s1 = []interface{}{1, 2, "你好", true}
	show(s1)
}
