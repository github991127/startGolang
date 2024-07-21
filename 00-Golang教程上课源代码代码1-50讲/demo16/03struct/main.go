/*

Go语言中的基础数据类型可以表示一些事物的基本属性，
但是当我们想表达一个事物的全部或部分属性时，
这时候再用单一的基本数据类型明显就无法满足需求了，
Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，
这种数据类型叫结构体，英文名称struct

注意：结构体首字母可以大写也可以小写，大写表示这个结构体是公有的，在其他的包里面 可以使用。小写表示这个结构体是私有的，只有这个包里面才能使用
*/

package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {

	//注意：在 Golang 中支持对结构体指针直接使用.来访问结构体的成员。p2.name = "张三" 其 实在底层是(*p2).name = "张三"

	var p2 = new(Person)
	p2.Name = "李四"
	p2.Age = 20
	p2.Sex = "男"
	(*p2).Name = "王五"
	fmt.Printf("值:%#v 类型:%T\n", p2, p2) //值:&main.Person{Name:"李四", Age:20, Sex:"男"} 类型:*main.Person

	var p3 = &Person{}
	p3.Name = "赵四"
	p3.Age = 23
	p3.Sex = "男"
	fmt.Printf("值:%#v 类型:%T\n", p3, p3) //值:&main.Person{Name:"赵四", Age:23, Sex:"男"} 类型:*main.Person

	var p4 = Person{
		Name: "哈哈",
		Age:  20,
		Sex:  "男",
	}
	fmt.Printf("值:%#v 类型:%T\n", p4, p4) //值:main.Person{Name:"哈哈", Age:20, Sex:"男"} 类型:main.Person

	var p5 = &Person{
		Name: "王麻子",
		Age:  20,
		Sex:  "男",
	}

	fmt.Printf("值:%#v 类型:%T\n", p5, p5) //值:&main.Person{Name:"王麻子", Age:20, Sex:"男"} 类型:*main.Person

	var p6 = &Person{
		Name: "王麻子",
	}
	fmt.Printf("值:%#v 类型:%T\n", p6, p6) //值:&main.Person{Name:"王麻子", Age:0, Sex:""} 类型:*main.Person

	var p7 = &Person{
		"张三",
		20,
		"男",
	}
	fmt.Printf("值:%#v 类型:%T\n", p7, p7) //值:&main.Person{Name:"张三", Age:20, Sex:"男"} 类型:*main.Person

}
