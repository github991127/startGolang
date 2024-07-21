package main

import "fmt"

func main() {
	//1、二维数组的定义
	// var arr = [3]int{1, 2, 3}  一维数组

	// var arr = [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// // fmt.Println(arr[0])
	// fmt.Println(arr[0][1])

	// var arr = [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// for _, v1 := range arr {
	// 	for _, v2 := range v1 {
	// 		fmt.Println(v2)
	// 	}
	// }

	// var arr = [3][2]string{
	// 	{"北京", "上海"},
	// 	{"广州", "深圳"},
	// 	{"成都", "重庆"},
	// }

	// for i := 0; i < len(arr); i++ {

	// 	for j := 0; j < len(arr[i]); j++ {
	// 		fmt.Println(arr[i][j])
	// 	}
	// }

	//定义二维数组的另一种方式
	var arr = [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
		{"成都", "重庆"},
		{"成都", "重庆"},
	}

	fmt.Println(arr)

}
