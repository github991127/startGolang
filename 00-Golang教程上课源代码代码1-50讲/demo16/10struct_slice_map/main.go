package main

import "fmt"

/*

结构体的字段类型可以是：基本数据类型、也可以是切片、Map 以及结构体

如果结构体的字段类型是: 指针，slice，和map的零值都是 nil ，即还没有分配空间

如果需要使用这样的字段，需要先make，才能使用.

*/

type Person struct {
	Name  string
	Age   int
	Hobby []string
	map1  map[string]string
}

func main() {

	var p Person
	p.Name = "张三"
	p.Age = 20
	p.Hobby = make([]string, 3, 6)
	p.Hobby[0] = "写代码"
	p.Hobby[1] = "打篮球"
	p.Hobby[2] = "睡觉"

	p.map1 = make(map[string]string)
	p.map1["address"] = "北京"
	p.map1["phone"] = "1324325325"

	fmt.Printf("%#v\n", p)

	fmt.Printf("%v", p.Hobby)

}
