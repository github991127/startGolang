package main

import "fmt"

func main() {
	/*
		值类型 ：改变变量副本值的时候，不会改变变量本身的值  (基本数据类型、数组)
		引用类型：改变变量副本值的时候，会改变变量本身的值 （切片、map）
	*/

	//map类型也是引用数据类型

	var userinfo1 = make(map[string]string)
	userinfo1["username"] = "张三"
	userinfo1["age"] = "20"
	userinfo2 := userinfo1

	userinfo2["username"] = "李四"
	fmt.Println(userinfo1)
	fmt.Println(userinfo2)

}
