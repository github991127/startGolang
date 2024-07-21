package main

import "fmt"

func main() {
	// var userinfo map[string]string
	// userinfo["username"] = "张三"
	// fmt.Println(userinfo)


	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// fmt.Println(userinfo)
	

	// var a []int
	// a[0] = 1
	// fmt.Println(a)

	// var a = make([]int, 4, 4)
	// a[0] = 1
	// fmt.Println(a)

	var a *int //指针也是引用数据类型
	*a = 100
	fmt.Println(*a)
}
