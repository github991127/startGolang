package main

import "fmt"

//一个结构体实现多个接口
type Animaler1 interface {
	SetName(string)
}

type Animaler2 interface {
	GetName() string
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

	var d1 Animaler1 = d //表示让Dog实现Animaler1这个接口
	var d2 Animaler2 = d //表示让Dog实现Animaler2这个接口

	d1.SetName("小花狗狗")
	fmt.Println(d2.GetName()) //小花狗狗

}
