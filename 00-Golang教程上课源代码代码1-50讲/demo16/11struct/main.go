package main

import "fmt"

/*
结构体嵌套
*/
type User struct {
	Username string
	Password string
	Address  Address //表示User结构体嵌套Address结构体
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

	fmt.Printf("%#v", u) //main.User{Username:"itying", Password:"1234567", Address:main.Address{Name:"张三", Phone:"15201671234", City:"北京"}}

}
