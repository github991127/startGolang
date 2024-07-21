package main

import "fmt"

/*
结构体值接收者实现接口：

值接收者： 如果结构体中的方法是值接收者，那么实例化后的结构体值类型和结构体指针类型都可以赋值给接口变量

*/
type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p Phone) start() { //值接收者
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {

	// 结构体值接收者例化后的结构体值类型和结构体指针类型都可以赋值给接口变量

	var p1 = Phone{
		Name: "小米手机",
	}
	var p2 Usber = p1 //表示让Phone实现Usb的接口
	p2.start()

	var p3 = &Phone{
		Name: "苹果手机",
	}
	var p4 Usber = p3 //表示让Phone实现Usb的接口
	p4.start()

}
