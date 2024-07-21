package main

import "fmt"

/*
结构体的匿名字段:

	结构体允许其成员字段在声明时没有字段名而只有类型，这种没有名字的字段就称为匿名字段

	匿名字段默认采用类型名作为字段名，结构体要求字段名称必须唯一，因此一个结构体中同种类型的匿名字段只能有一个
*/
type Person struct {
	string
	int
}

func main() {

	p := Person{
		"张三",
		20,
	}
	fmt.Println(p)
}
