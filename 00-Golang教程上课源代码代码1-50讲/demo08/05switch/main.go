package main

import "fmt"

func main() {
	/*

		switch expression {
		case condition:


		}
	*/

	// 练习：判断文件类型,如果后缀名是.html 输入 text/html, 如果后缀名.css 输出 text/css ,如果 后缀名是.js 输出 text/javascript

	// var ext = ".sss"
	// if ext == ".html" {
	// 	fmt.Println("text/html")
	// } else if ext == ".css" {
	// 	fmt.Println("text/css")
	// } else if ext == ".js" {
	// 	fmt.Println("text/javascript")
	// } else {
	// 	fmt.Println("找不到此后缀")
	// }

	// 1、switch case的基本使用
	// var extname = ".css"
	// switch extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// case ".js":
	// 	fmt.Println("text/javascript")
	// 	break
	// default:
	// 	fmt.Println("找不到此后缀")
	// }

	// 2、switch case的另一种写法

	// switch extname := ".html"; extname {
	// case ".html":
	// 	fmt.Println("text/html")
	// 	break
	// case ".css":
	// 	fmt.Println("text/css")
	// 	break
	// case ".js":
	// 	fmt.Println("text/javascript")
	// 	break
	// default:
	// 	fmt.Println("找不到此后缀")
	// }
	// fmt.Println(extname)  //undefined: extname

	// 3、一个分支可以有多个值，多个 case 值中间使用英文逗号分隔

	// 判断一个数是不是偶数

	// var n = 8
	// switch n {
	// case 1, 3, 5, 7, 9:
	// 	fmt.Println("奇数")
	// 	break //golang中break可以写也可以不写
	// case 2, 4, 6, 8, 10:
	// 	fmt.Println("偶数")
	// 	break
	// }

	// var score = "D" //ABC及格  D不及格
	// switch score {
	// case "A", "B", "C":
	// 	fmt.Println("及格")
	// case "D":
	// 	fmt.Println("不及格")
	// }

	//ABC及格  D不及格
	// switch score := "D"; score {
	// case "A", "B", "C":
	// 	fmt.Println("及格")
	// case "D":
	// 	fmt.Println("不及格")
	// }

	//4、分支还可以使用表达式，这时候 switch 语句后面不需要再跟判断变量。例如 

	// var age = 18
	// switch {
	// case age < 24:
	// 	fmt.Println("好好学习")
	// case age >= 24 && age <= 60:
	// 	fmt.Println("好好赚钱")
	// case age > 60:
	// 	fmt.Println("注意身体")
	// default:
	// 	fmt.Println("输入错误")
	// }

	//5、 switch 的穿透 fallthrought

	//fallthrough`语法可以执行满足条件的 case 的下一个 case，是为了兼容 C 语言中的 case 设计 的

	var age = 30
	switch {
	case age < 24:
		fmt.Println("好好学习")
	case age >= 24 && age <= 60:
		fmt.Println("好好赚钱")
		fallthrough
	case age > 60:
		fmt.Println("注意身体")
	default:
		fmt.Println("输入错误")
	}
}
