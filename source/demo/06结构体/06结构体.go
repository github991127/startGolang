package main

import (
	"encoding/json"
	"fmt"
)

type Person struct { //大写开头的结构体名表示该结构体为公共类型
	Name string //`json:"id"` // 指定json的key名，默认是结构体的字段名
	Age  int
	bool // 1匿名字段：只有类型，采用类型名作为字段名，一个结构体只能有一个同类型的匿名字段
	//s1 []string // 2指针，slice，map字段：零值都是 nil ，还没有分配空间，需要先make才能使用.
	//map1 map[string]string
	//A1 Address// 3结构体嵌套：一个结构体可以包含另一个结构体，类似子类继承
	//Address // 结构体嵌套支持匿名，当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找
	//*Address //结构体嵌套支持指针
}

func (p Person) PrintInfo() {
	fmt.Printf("Name:%s Age:%d\n", p.Name, p.Age)
}

// 结构体的声明与初始化
func f1() {
	p1 := Person{
		Name: "Alice",
		Age:  25,
	}
	p1.PrintInfo()                        //Name:Alice Age:25
	fmt.Printf("p1值:%#v 类型:%T\n", p1, p1) //p1值:main.Person{Name:"Alice", Age:25} 类型:main.Person

	p2 := new(Person)
	p2.Name = "Bob"
	p2.Age = 30
	p2.PrintInfo()                        //Name:Bob Age:30
	fmt.Printf("p2值:%#v 类型:%T\n", p2, p2) //p2值:&main.Person{Name:"Bob", Age:30} 类型:*main.Person

}

// 结构体与json的互相转换
func f2() {
	p1 := Person{
		Name: "Alice",
		Age:  25,
	}
	// 结构体转json（须字段公有）
	jsonByte, _ := json.Marshal(p1)
	jsonStr := string(jsonByte)
	fmt.Printf("%v\n", jsonStr)
	// json转结构体
	jsonStr = `{"Name":"Bob","Age":30}` //``取消转义
	var p2 Person
	err := json.Unmarshal([]byte(jsonStr), &p2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", p2)

}

func main() {
	//f1() // 结构体的声明与初始化
	f2() // 结构体与json的互相转换
}
