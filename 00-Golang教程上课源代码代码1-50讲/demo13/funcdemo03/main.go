package main

import "fmt"

//int类型升序排序
func sortIntAsc(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}

	return slice
}

//int类型降序排序  切片是引用数据类型
func sortIntDesc(slice []int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}

}

func main() {

	//案例1：把前面讲的选择排序封装成方法，实现整型切片的升序降序排序排列

	// var sliceA = []int{12, 34, 37, 35, 556, 36, 2}
	// arr := sortIntAsc(sliceA)
	// fmt.Println(arr)

	// var sliceB = []int{1, 34, 4, 35, 6, 36, 2}
	// fmt.Println(sortIntAsc(sliceB))

	var sliceC = []int{1, 34, 4, 36, 126, 36, 2}
	sortIntDesc(sliceC)
	fmt.Println(sliceC)

}
