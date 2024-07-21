/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

//接口是一个规范
type Usber interface {
	start()
	stop()
}

//如果接口里面有方法的话，必要要通过结构体或者通过自定义类型实现这个接口

type Phone struct {
	Name string
}

//手机要实现usb接口的话必须得实现usb接口中的所有方法
func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}

func (p Camera) stop() {
	fmt.Println("相机关机")
}

func (p Camera) run() {
	fmt.Println("run")
}

func main() {

	p := Phone{
		Name: "华为手机",
	}
	// p.start1()

	var p1 Usber //golang中接口就是一个数据类型
	p1 = p       //表示手机实现Usb接口
	p1.stop()

	c := Camera{}
	var c1 Usber = c //表示相机实现了Usb接口
	c1.start()
	// c1.run()  //错误：type Usber has no field or method run)
	c.run() //run

}
