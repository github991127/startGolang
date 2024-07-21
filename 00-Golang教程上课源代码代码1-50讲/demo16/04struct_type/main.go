/*
	值类型 ： 改变变量副本值的时候，不会改变变量本身的值 (数组、基本数据类型、结构体)
	引用类型：改变变量副本值的时候，会改变变量本身的值  （切片、map）
*/

package main

import "fmt"

type Person struct {
	Name string
	Age  int
	Sex  string
}

func main() {
	var p1 = Person{
		Name: "哈哈",
		Age:  20,
		Sex:  "男",
	}

	p2 := p1
	p2.Name = "李四"
	fmt.Printf("%#v\n", p1) //main.Person{Name:"哈哈", Age:20, Sex:"男"}
	fmt.Printf("%#v", p2)   //main.Person{Name:"李四", Age:20, Sex:"男"}

}
