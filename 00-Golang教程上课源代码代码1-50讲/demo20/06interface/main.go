package main

import "fmt"

type Usber interface {
	start()
	stop()
}

//电脑
type Computer struct {
}

func (c Computer) work(usb Usber) {
	//要判断usb的类型
	if _, ok := usb.(Phone); ok { //类型断言
		usb.start()
	} else {
		usb.stop()
	}

}

//手机
type Phone struct {
	Name string
}

func (p Phone) start() {
	fmt.Println(p.Name, "启动")
}

func (p Phone) stop() {
	fmt.Println(p.Name, "关机")
}

//照相机
type Camera struct {
}

func (p Camera) start() {
	fmt.Println("相机启动")
}
func (p Camera) stop() {
	fmt.Println("相机关机")
}

func main() {

	var computer = Computer{}
	var phone = Phone{
		Name: "小米手机",
	}
	var camera = Camera{}

	computer.work(phone)
	computer.work(camera)

}
