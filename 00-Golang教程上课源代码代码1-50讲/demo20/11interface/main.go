package main

import "fmt"

//接口嵌套
type Ainterface interface {
	SetName(string)
}

type Binterface interface {
	GetName() string
}

type Animaler interface { //接口的嵌套
	Ainterface
	Binterface
}

type Dog struct {
	Name string
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d Dog) GetName() string {
	return d.Name
}

func main() {

	//Dog实现Animal的接口
	d := &Dog{
		Name: "小黑",
	}

	var d1 Animaler = d //表示让Dog实现Animaler这个接口
	d1.SetName("小花狗狗")
	fmt.Println(d1.GetName()) //小花狗狗

}
