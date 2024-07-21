package main

import "fmt"

// 接口定义
type Usber interface {
	start()
	stop()
}

// 总线接口实现
type Computer struct {
}

func (c Computer) work(usb Usber) {
	usb.start()
	usb.stop()
	//断言，判断接口类型
	if _, ok := usb.(Phone); ok { //类型断言
		//usb.start()
		fmt.Println("phone start")
	} else {
		//usb.stop()
		fmt.Println("others stop")
	}

}

// 接口实例，手机
type Phone struct {
	Name string
}

// 实现接口方法
func (p Phone) start() {
	fmt.Println(p.Name, "接口方法1，启动")
}
func (p Phone) stop() {
	fmt.Println(p.Name, "接口方法2，关闭")
}

//如果结构体中的方法是指针接收者，则需要使用指针调用
//func (p *Phone) start() { //指针接收者
//	fmt.Println(p.Name, "启动")
//}

//var phone1 = &Phone{
//	Name: "小米",
//}

func main() {
	var c = Computer{}
	var p = Phone{"小米"}
	c.work(p) // 多态：不同实例调用接口方法
}
