package main

import "fmt"

func main() {
	/*
		什么叫做冒泡排序？
			概念:从头到尾,比较相邻的两个元素的大小,如果符合交换条件,交换两个元素的位置。
			特点:每一轮比较中,都会选出一个最大的数，放在正确的位置。

			第一轮:
			9 8 7 6 5 4

			8 9 7 6 5 4
			8 7 9 6 5 4
			8 7 6 9 5 4
			8 7 6 5 9 4
			8 7 6 5 4 9    9放在了正确的位置


			第二轮:
			8 7 6 5 4

			7 8 6 5 4
			7 6 8 5 4
			7 6 5 8 4
			7 6 5 4 8     8放在了正确的位置


			第三轮:
			7 6 5 4

			6 7 5 4
			6 5 7 4
			6 5 4 7 	  7放在了正确的位置

			第四轮:
			6 5 4

			5 6 4
			5 4 6         6放在了正确的位置

			第五轮:
			5 4

			4 5
	*/

	//冒泡排序从小到大
	// var numSlice = []int{9, 6, 5, 4, 8}
	// for i := 0; i < len(numSlice); i++ {
	// 	for j := 0; j < len(numSlice)-1-i; j++ {
	// 		if numSlice[j] > numSlice[j+1] {
	// 			temp := numSlice[j]
	// 			numSlice[j] = numSlice[j+1]
	// 			numSlice[j+1] = temp
	// 		}

	// 	}
	// }

	// fmt.Println(numSlice)

	//冒泡排序从大到下
	var numSlice = []int{9, 6, 5, 4, 8}
	for i := 0; i < len(numSlice); i++ {
		for j := 0; j < len(numSlice)-1-i; j++ {
			if numSlice[j] < numSlice[j+1] {
				temp := numSlice[j]
				numSlice[j] = numSlice[j+1]
				numSlice[j+1] = temp
			}

		}
	}
	fmt.Println(numSlice)

}
