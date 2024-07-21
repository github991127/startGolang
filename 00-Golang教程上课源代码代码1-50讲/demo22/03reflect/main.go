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
	// v.Name() //类型名称 ,种类（Kind）就是指底层的类型
	// v.Kind() //种类
	fmt.Printf("类型:%v 类型名称:%v 类型种类:%v \n", v, v.Name(), v.Kind())
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
	reflectFn(&h) //*int 类型名称: 类型种类:ptr

	var i = [3]int{1, 2, 3}
	reflectFn(i) //[3]int 类型名称: 类型种类:array

	var j = []int{11, 22, 33}
	reflectFn(j) //[]int 类型名称: 类型种类:slice

}
