/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"`
	Sno    string `json:"sno"`
}

func main() {
	var s1 = Student{
		ID:     1,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Println("aa")

	var s, _ = json.Marshal(s1)
	jsonStr := string(s)
	fmt.Println(jsonStr)
}
