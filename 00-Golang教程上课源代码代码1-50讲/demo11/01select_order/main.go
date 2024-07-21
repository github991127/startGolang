/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import "fmt"

func main() {
	//回顾练习：从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	// var arr = [...]int{1, 3, 5, 7, 8}
	// for i := 0; i < len(arr); i++ {
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[i]+arr[j] == 8 {
	// 			fmt.Printf("(%v,%v)", i, j)
	// 		}
	// 	}
	// }

	/*
		选择排序：进行从小到大排序
		概念: 通过比较 首先选出最小的数放在第一个位置上，然后在其余的数中选出次小数放在第二个位置上,依此类推,直到所有的数成为有序序列。
		[9, 8, 7, 6, 5, 4]
	*/

	/*

		第一轮:
			9 8 7 6 5 4
			8 9 7 6 5 4
			7 9 8 6 5 4
			6 9 8 7 5 4
			5 9 8 7 6 4
			4 9 8 7 6 5    4放在了正确的位置
		第二轮:
			  9 8 7 6 5
			  8 9 7 6 5
			  7 9 8 6 5
			  6 9 8 7 5
			  5 9 8 7 6    5放在了正确的位置

		第三轮:
				9 8 7 6
				8 9 7 6
				7 9 8 6
				6 9 8 7    6放在了正确的位置

		第四轮:
				  9 8 7
				  8 9 7
				  7 9 8    7放在了正确的位置
		第五轮:
					9 8
					8 9    8放在了正确的位置

	*/

	//numSlice[0] = 4
	//numSlice[1] = 5
	//numSlice[2] = 6
	//numSlice[3] = 7
	//numSlice[4] = 8
	//numSlice[5] = 8

	//从小到大的排序
	// var numSlice = []int{9, 8, 7, 6, 5, 4}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := i + 1; j < len(numSlice); j++ {
	// 		if numSlice[i] > numSlice[j] {
	// 			temp := numSlice[i]
	// 			numSlice[i] = numSlice[j]
	// 			numSlice[j] = temp
	// 		}

	// 	}
	// }
	// fmt.Println(numSlice)

	//从大到小的排序
	var numSlice = []int{6, 5, 4, 9, 8, 7}
	for i := 0; i < len(numSlice); i++ {
		for j := i + 1; j < len(numSlice); j++ {
			if numSlice[i] < numSlice[j] {
				temp := numSlice[i]
				numSlice[i] = numSlice[j]
				numSlice[j] = temp
			}

		}
	}
	fmt.Println(numSlice)
}
