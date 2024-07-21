/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

func getUserinfo() (string, int) {
	return "zhangsan", 10
}

var g = "全局变量"

// gl := "全局变量 _端变量声明法  错误写法"

func main() {
	fmt.Println(g)
	// fmt.Println(gl)
	/*
	   1、 var 声明变量
	   	var 变量名称 类型

	   	变量名称的命名：Go 语言变量名由字母、数字、下划线组成，其中首个字符不能为数字。Go 语言中关键字 和保留字都不能用作变量名。
	*/

	// var username string
	// fmt.Println(username)   //变量声明后没有初始化的话  值为空

	// var a1="张三"
	// fmt.Println(a1)

	// var 1a string   //错误：go变量的首个字符不能为数字
	// fmt.Println(1a)

	// var m_a="李四"   //正确 不推荐
	// fmt.Println(m_a)

	// var if="李四"  //错误  Go 语言中关键字 和保留字都不能用作变量名。
	// fmt.Println(if)

	//go语言变量的定义以及初始化

	//第一种初始化变量的方式
	// var username string
	// username = "张三"
	// fmt.Println(username)

	//第二种初始化变量的方式
	// var username string = "张三"
	// fmt.Println(username)

	//第三种初始化变量的方式
	// var username = "张三"
	// fmt.Println(username)

	/*
		Go 语言中的变量需要声明后才能使用，同一作用域内不支持重复声明
		var username = "张三"
		var age = 20
		var sex="男"
		fmt.Println(username,age,sex)
		var username = "李四"
	*/

	/*
		var username = "张三"
		var age = 20
		var sex="男"

		username = "李四"   //赋值

		fmt.Println(username,age,sex)  //输出：李四 20 男
	*/

	/*

		2、 一次定义多个变量

			var 变量名称, 变量名称 类型

			var (
				变量名称 类型
				变量名称 类型
			)
	*/

	/*
		var a1, a2 string
		a1 = "aaa"
		a2 = "aaaaaaaaaaa"
		fmt.Println(a1, a2)
	*/

	/*
		var a1, a2 string
		a1 = "aaa"
		a2 = 123 //错误写法  go语言中定义的类型是string类型 赋值的时候必须赋值string类型的数据
		fmt.Println(a1, a2)
	*/

	/*
		var (
			username string
			age      int
			sex      string
		)

		username = "张三"
		age = 20
		sex = "男"

		fmt.Println(username, age, sex)
	*/

	/*
		var (
			username = "张三"
			age      = 20
		)
		fmt.Println(username, age)
	*/

	/*
		3、短变量声明法 在函数内部，可以使用更简略的 := 方式声明并初始化变量。

		注意：短变量只能用于声明局部变量，不能用于全局变量的声明
	*/

	// var username = "张三"
	/*
		username := "张三"

		fmt.Println(username)

		fmt.Printf("类型:%T", username)
	*/

	// 使用端变量一次声明多个变量，并初始化变量

	// a, b, c := 12, 13, 20
	// fmt.Println(a, b, c)

	// a, b, c := 12, 13, "C"
	// fmt.Println(a, b, c)
	// fmt.Printf("a类型:%T b类型:%T c类型:%T", a, b, c)

	/*
		4、匿名变量 在使用多重赋值时，如果想要忽略某个值，可以使用匿名变量（anonymous variable）。

		匿 名变量用一个下划线_表示，例如

		func getUserinfo() (string, int) {
			return "zhangsan", 10
		}
	*/
	// var username, age = getUserinfo()
	// fmt.Println(username, age) //zhangsan 10

	var username, _ = getUserinfo()
	fmt.Println(username) //zhangsan
	//匿名变量不占用命名空间，不会分配内存，所以匿名变量之间不存在重复声明
	var _, age = getUserinfo()
	fmt.Println(age) //10
}
