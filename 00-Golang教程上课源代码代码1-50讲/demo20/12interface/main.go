/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

type Address struct {
	Name  string
	Phone int
}

// Golang中空接口和类型断言使用细节
func main() {

	var userinfo = make(map[string]interface{})
	userinfo["username"] = "张三"
	userinfo["age"] = 20
	userinfo["hobby"] = []string{"睡觉", "吃饭"}

	fmt.Println(userinfo["age"])
	fmt.Println(userinfo["hobby"])

	// fmt.Println(userinfo["hobby"][1]) //interface {} does not support indexing

	var address = Address{
		Name:  "李四",
		Phone: 1521242141,
	}
	fmt.Println(address.Name) //李四

	userinfo["address"] = address

	fmt.Println(userinfo["address"]) //{李四 1521242141}

	// var name = userinfo["address"].Name //type interface {} is interface with no methods
	// fmt.Println(name)

	hobby2, _ := userinfo["hobby"].([]string)

	fmt.Println(hobby2[1]) //吃饭

	address2, _ := userinfo["address"].(Address)
	fmt.Println(address2.Name, address2.Phone) //李四 1521242141

}
