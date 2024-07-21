package main

import "fmt"

func main() {
	//1、数组的长度是类型的一部分
	// var arr1 [3]int
	// var arr2 [4]int
	// var strArr [3]string

	// fmt.Printf("arr1:%T  arr2:%T strArr:%T", arr1, arr2, strArr)

	//2、数组的初始化 第一种方法

	// var arr1 [3]int

	// arr1[0] = 23
	// arr1[1] = 10
	// arr1[2] = 24
	// fmt.Println(arr1)

	// var strArr [3]string
	// strArr[0] = "php"
	// strArr[1] = "java"
	// strArr[2] = "golang"
	// // strArr[3] = "js" //out of bounds for 3-element array
	// fmt.Println(strArr)

	//2、数组的初始化 第二种方法

	// var arr1 = [3]int{23, 34, 5}
	// fmt.Println(arr1)

	// var arr1 = [3]string{"php", "nodejs", "golnag"}
	// fmt.Println(arr1)

	// arr1 := [3]string{"php", "nodejs", "golnag"}

	// fmt.Println(arr1)

	//3、数组的初始化 第三种方法  一般情况下我们可以让编译器 根据初始值的个数自行推断数组的长度

	// var arr1 = [...]int{12, 2324, 32435, 353, 535, 3535}

	// fmt.Println(arr1)

	// fmt.Println(len(arr1)) //len()打印数组的长度

	//注意数组的长度

	// arr1 := [...]string{"php", "nodejs", "golnag", "js"}
	// arr1[4] = "java" //: invalid array index 4 (out of bounds for 4-element array)
	// fmt.Println(arr1)

	//改变数组里面的值
	// arr1 := [...]string{"php", "nodejs", "golnag", "js"}
	// arr1[0] = "phper"
	// fmt.Println(arr1)

	// 4、数组的初始化 第四种方法   我们还可以使用指定索引值的方式来初始化数组

	// arr := [...]int{0: 1, 1: 10, 2: 20, 5: 50}

	// fmt.Println(len(arr)) //6

	// fmt.Println(arr) //[1 10 20 0 0 50]

	//5、数组的循环遍历 for  for range

	// var arr = [3]int{23, 34, 5}
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }

	// arr1 := [...]string{"php", "nodejs", "golnag", "js"}
	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Println(arr1[i])
	// }

	arr1 := [...]string{"php", "nodejs", "golnag", "js"}

	for k, v := range arr1 {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

}
