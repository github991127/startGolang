package main

import "fmt"

//定义一个Animal的接口，Animal中定义两个方法，分别是SetName和GetName。分别让Dog结构体和Cat结构体实现这个方法

type Animaler interface {
	SetName(string)
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

type Cat struct {
	Name string
}

func (c *Cat) SetName(name string) {
	c.Name = name
}
func (c Cat) GetName() string {
	return c.Name
}

func main() {

	//Dog实现Animal的接口
	d := &Dog{
		Name: "小黑",
	}
	var d1 Animaler = d
	fmt.Println(d1.GetName())
	d1.SetName("阿奇")
	fmt.Println(d1.GetName())

	//Cat实现Animal的接口
	c := &Cat{
		Name: "小花",
	}
	var c1 Animaler = c
	fmt.Println(c1.GetName())
}
