package main

import "fmt"

/*
嵌套匿名结构体
*/
type User struct {
	Username string
	Password string
	AddTime  string
	Address
}
type Address struct {
	Name    string
	Phone   string
	City    string
	AddTime string
}

func main() {
	var u User
	u.Username = "itying"
	u.Password = "1234567"
	u.Address.Name = "张三"
	u.Address.Phone = "15201671234"
	u.Address.City = "北京"

	u.City = "上海" //当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

	u.AddTime = "2020-05-1"

	u.Address.AddTime = "2020-06-1"

	fmt.Printf("%#v\n", u)

}
