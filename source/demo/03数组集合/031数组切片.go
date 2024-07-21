package main

import (
	"fmt"
	"sort"
)

// 一维数组
func f1() {
	//初始化
	// var arr1 [3]int
	// arr1[0] = 20
	// arr1[1] = 30
	// arr1[2] = 50
	//arr1 :=  [3]int{20, 30, 50}
	//arr1 :=  [...]int{20, 30, 50}//自行推断数组的长度
	arr1 := [...]int{0: 20, 2: 30, 5: 50} //指定索引和值
	fmt.Println(arr1)
	fmt.Printf("arr1:%T\n", arr1)
	arr2 := [3]string{"php", "nodejs", "golnag"}
	fmt.Println(arr2)
	fmt.Printf("arr2:%T\n", arr2)

	//访问元素
	for k, v := range arr2 {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	//数组值传递，切片引用传递
	var arr3 = [...]int{1, 2, 3}
	arr4 := arr3
	arr4[0] = 11
	fmt.Println(arr3) //[1 2 3]
	fmt.Println(arr4) //[11 2 3]

	var arr5 = []int{1, 2, 3}
	arr6 := arr5
	arr6[0] = 11
	fmt.Println(arr5) //[11 2 3]
	fmt.Println(arr6) //[11 2 3]

}

// 二维数组
func f2() {
	//初始化
	var arr = [...][2]string{ // 行数自动推断
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
		{"成都", "重庆"},
		{"成都", "重庆"},
	}
	fmt.Println(arr)
	fmt.Println(arr[0])
	fmt.Println(arr[0][1])

	//访问元素
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("arr[%d][%d]=%s\n", i, j, arr[i][j])
		}
	}
	for _, v1 := range arr {
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}

}

// 切片
func f3() {
	//初始化
	//make()函数创建一个切片  make([]T, size, cap)
	var arr = make([]int, 4, 8)
	fmt.Printf("长度%d 容量%d\n", len(arr), cap(arr))

	var arr1 []int
	var arr2 = []int{1: 20, 2: 30, 5: 50}
	fmt.Println(arr1, len(arr1))
	fmt.Println(arr2, len(arr2)) //[]
	fmt.Println(arr1 == nil)     //true golnag中申明切片以后 切片的默认值就是nil

	//可以基于数组或切片创建新的切片
	//容量=底层数组元素个数-第一个元素的索引
	b := arr2[1:3]
	fmt.Printf("长度%d 容量%d\n", len(b), cap(b)) //长度2 容量5
	c := arr2[:3]
	fmt.Printf("长度%d 容量%d\n", len(c), cap(c)) //长度3 容量6
	d := c[2:]
	fmt.Printf("长度%d 容量%d\n", len(d), cap(d)) //长度2 容量4

	//增删
	arr = append(arr, 10)             //增加元素
	arr = append(arr, arr2...)        //增加切片
	arr = append(arr[:2], arr[3:]...) //删除索引为2的元素
	fmt.Printf("长度%d 容量%d\n", len(arr), cap(arr))

	//切片复制
	sliceA := []int{1, 2, 3, 4}
	sliceB := make([]int, 4, 4)
	copy(sliceB, sliceA) //切片是引用类型，为避免一损俱损，需要使用copy()函数进行复制改为值传递
	sliceB[0] = 111
	fmt.Println(sliceA) //[1 2 3 4]
	fmt.Println(sliceB) //[111 2 3 4]

	//字符串字节切片
	s1 := "big"
	byteStr := []byte(s1)
	byteStr[0] = 'p'
	fmt.Println(string(byteStr))

	//字符串字符切片
	s2 := "你好golang"
	runeStr := []rune(s2)
	fmt.Println(runeStr) //[20320 22909 103 111 108 97 110 103]
	runeStr[0] = '大'
	fmt.Println(string(runeStr))

}

func f4() {
	//升序排序
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	sort.Ints(arr)
	// sort.Float64s(arr)
	// sort.Strings(arr)
	fmt.Println(arr)

	//降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	fmt.Println(arr)
}
func main() {
	//f1() // 一维数组
	//f2() // 二维数组
	//f3() // 切片
	f4() // 排序

}
