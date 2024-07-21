package main

import (
	"encoding/json"
	"fmt"
)

//结构体标签
type Student struct {
	Id     int    `json:"id"`
	Gender string `json:"gender"`
	Name   string `json:"name"` //私有属性不能被json包访问
	Sno    string `json:"sno"`
}

func main() {
	var s1 = Student{
		Id:     12,
		Gender: "男",
		Name:   "李四",
		Sno:    "s0001",
	}
	fmt.Printf("%#v\n", s1) //main.Student{Id:12, Gender:"男", Name:"李四", Sno:"s0001"}

	jsonByte, _ := json.Marshal(s1)
	jsonStr := string(jsonByte)
	fmt.Printf("%v", jsonStr) //{"id":12,"gender":"男","name":"李四","sno":"s0001"}
}
