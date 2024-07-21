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

// 反射获取任意变量的类型
func reflectFn(x interface{}) {
	v := reflect.TypeOf(x)
	// v.Name() //类型名称 ,种类（Kind）就是指底层的类型
	// v.Kind() //种类
	fmt.Printf("类型:%v 类型名称:%v 类型种类:%v \n", v, v.Name(), v.Kind())
}

// 反射获取任意变量的原始值
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	// v.Kind() //获取种类
	kind := v.Kind()

	switch kind {
	case reflect.Int64:
		fmt.Printf("int类型的原始值%v,计算后的值是%v \n", v.Int(), v.Int()+10)
	case reflect.Float32:
		fmt.Printf("Float32类型的原始值%v\n", v.Float())
	case reflect.Float64:
		fmt.Printf("Float64类型的原始值%v\n", v.Float())
	case reflect.String:
		fmt.Printf("string类型的原始值%v\n", v.String())
	default:
		fmt.Printf("还没有判断这个类型\n")
	}

}

// 反射获取任意变量的类型，测试用例
func f1() {
	a := 10
	b := 23.4
	c := true
	d := "你好golang"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

	var e myInt = 34
	var f = Person{
		Name: "张三",
		Age:  20,
	}
	reflectFn(e)
	reflectFn(f)

	var h = 25
	reflectFn(&h) //*int 类型名称: 类型种类:ptr

	var i = [3]int{1, 2, 3}
	reflectFn(i) //[3]int 类型名称: 类型种类:array

	var j = []int{11, 22, 33}
	reflectFn(j) //[]int 类型名称: 类型种类:slice
}

// 反射获取任意变量的原始值，测试用例
func f2() {
	var a int64 = 100
	var b float32 = 12.3
	var c string = "你好golang"
	reflectValue(a)
	reflectValue(b)
	reflectValue(c)
}

func main() {
	f1()
	f2()
}
