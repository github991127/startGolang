package main

import "fmt"

func main() {
	// 1、请求出一个数组里面元素的和以及这些元素的平均值。分别用for和for-range实现

	// var arr = [...]int{12, 23, 45, 67, 2, 5}
	// var sum = 0
	// for i := 0; i < len(arr); i++ {
	// 	sum += arr[i]
	// }
	// fmt.Printf("arr数组元素的和是:%v 平均值：%.2f", sum, float64(sum)/float64(len(arr)))

	/*
		2、请求出一个数组的最大值，并得到对应的下标

				思路
					1、 声明一个数组 var intArr = [...]int {1, -1, 12, 65, 11}
					2、假定第一个元素就是最大值，下标就 0
					3、然后从第二个元素开始循环比较，如果发现有更大，则交换
	*/

	// var arr = [...]int{1, -1, 12, 65, 11}
	// var max = arr[0]
	// var index = 0
	// for i := 0; i < len(arr); i++ {
	// 	if max < arr[i] {
	// 		max = arr[i]
	// 		index = i
	// 	}
	// }
	// fmt.Printf("最大值:%v  最大值对应的索引值:%v", max, index)

	//3、从数组[1, 3, 5, 7, 8]中找出和为8的两个元素的下标分别为(0,3)和(1,2)。

	var arr = [...]int{1, 3, 5, 7, 8}

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == 8 {
				fmt.Printf("(%v,%v)", i, j)
			}
		}
	}

}
