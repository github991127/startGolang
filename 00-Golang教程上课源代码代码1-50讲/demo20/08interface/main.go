package main

import "fmt"

/*
结构体指针接收者实现接口：

指针接收者： 如果结构体中的方法是指针接收者，那么实例化后结构体指针类型都可以赋值给接口变量， 结构体值类型没法赋值给接口变量。

*/
type Usber interface {
	start()
	stop()
}

type Phone struct {
	Name string
}

func (p *Phone) start() { //指针接收者
	fmt.Println(p.Name, "启动")
}

func (p *Phone) stop() {
	fmt.Println(p.Name, "关机")
}

func main() {

	/*
	   错误写法
	   	var phone1 = Phone{
	   		Name: "小米",
	   	}
	   	var p1 Usber = phone1 // Phone does not implement Usber (start method has pointer receiver
	   	p1.start()
	*/

	var phone1 = &Phone{
		Name: "小米",
	}
	var p1 Usber = phone1
	p1.start()
}
