package main

import "fmt"

/*
嵌套匿名结构体
*/
type User struct {
	Username string
	Password string
	Address
}
type Address struct {
	Name  string
	Phone string
	City  string
}

func main() {
	var u User
	u.Username = "itying"
	u.Password = "1234567"
	u.Address.Name = "张三"
	u.Address.Phone = "15201671234"
	u.Address.City = "北京"

	u.City = "上海" //当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。

	fmt.Printf("%#v\n", u) //main.User{Username:"itying", Password:"1234567", Address:main.Address{Name:"张三", Phone:"15201671234", City:"上海"}}

	fmt.Println(u.Address.Phone)
	fmt.Println(u.Phone)
}
