/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"fmt"
	"reflect"
)

//student结构体
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Score int    `json:"score"`
}

func (s Student) GetInfo() string {
	var str = fmt.Sprintf("姓名:%v 年龄:%v 成绩:%v", s.Name, s.Age, s.Score)
	return str
}

//反射修改结构体属性
func reflectChangeStruct(s interface{}) {
	t := reflect.TypeOf(s)
	v := reflect.ValueOf(s)

	if t.Kind() != reflect.Ptr {
		fmt.Println("传入的不是结构体指针类型")
		return
	} else if t.Elem().Kind() != reflect.Struct {

		fmt.Println("传入的不是结构体指针类型")
		return
	}
	//修改结构体属性的值
	name := v.Elem().FieldByName("Name")
	name.SetString("小李")

	age := v.Elem().FieldByName("Age")
	age.SetInt(22)

}
func main() {
	stu1 := Student{
		Name:  "小明",
		Age:   15,
		Score: 98,
	}
	// PrintStructField(stu1)
	reflectChangeStruct(&stu1)

	fmt.Printf("%#v\n", stu1) //main.Student{Name:"小李", Age:22, Score:98}

}
