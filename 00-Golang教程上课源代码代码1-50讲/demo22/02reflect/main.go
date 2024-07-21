package main

import (
	"fmt"
	"reflect"
)

type myInt int
type Person struct {
	Name string
	Age  int
}

//反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Println(v)
}
func main() {
	a := 10         //int
	b := 23.4       //float64
	c := true       //bool
	d := "你好golang" //string
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	var f = Person{
		Name: "张三",
		Age:  20,
	}
	reflectFn(e) //main.myInt
	reflectFn(f) //main.Person

	var h = 25
	reflectFn(&h) //*int

}
