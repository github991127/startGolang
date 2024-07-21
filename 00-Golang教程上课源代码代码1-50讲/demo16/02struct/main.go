/*

Go语言中的基础数据类型可以表示一些事物的基本属性，
但是当我们想表达一个事物的全部或部分属性时，
这时候再用单一的基本数据类型明显就无法满足需求了，
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，
这种数据类型叫结构体，英文名称struct


*/
package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  string
}

func main() {

	var p1 Person //实例化Person结构体
	p1.name = "张三"
	p1.sex = "男"
	p1.age = 20
	fmt.Printf("值:%v 类型:%T\n", p1, p1) //值:{张三 20 男} 类型:main.Person
	fmt.Printf("值:%#v 类型:%T", p1, p1)  //值:main.Person{name:"张三", age:20, sex:"男"} 类型:main.Person

}
