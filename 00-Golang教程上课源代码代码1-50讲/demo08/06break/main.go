package main

import "fmt"

func main() {
	
	/*
		Go 语言中 break 语句用于以下几个方面：
			• 用于循环语句中跳出循环，并开始执行循环之后的语句。
			• break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。
			• 在多重循环中，可以用标号 label 标出想 break 的循环。
	*/

	//1、表示当i=2的时候跳出当前循环
	// for i := 1; i <= 10; i++ {
	// 	if i == 2 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }
	// fmt.Println("继续执行")

	/*
		i=0

			i=0 j=0
			i=0 j=1
			i=0 j=2

		i=1
			i=1 j=0
			i=1 j=1
			i=1 j=2
	*/

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 3 {
	// 			break
	// 		}
	// 		fmt.Printf("i=%v j=%v\n", i, j)
	// 	}
	// }

	//2、 break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。

	// var extname = ".html"
	// switch extname {
	// case ".html":
	// 	fmt.Println("text/html")
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

	//3、在多重循环中，可以用标号 label 标出想 break 的循环。

	/*
		i=0 j=0
		i=0 j=1
		i=0 j=2
	*/
lable1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break lable1
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}
}
