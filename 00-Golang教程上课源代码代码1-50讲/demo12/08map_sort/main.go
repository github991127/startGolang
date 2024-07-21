package main

import (
	"fmt"
	"sort"
)

func main() {
	//map的排序
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90
	map1[11] = 100
	map1[2] = 13

	// fmt.Println(map1)

	/*
		1 13
		4 56
		10 100
		8 90
	*/
	// for key, val := range map1 {
	// 	fmt.Println(key, val)
	// }

	//需求：按照key升序输出map的key=>value   (签名算法)

	//1、把map的key放在切片里面
	var keySlice []int
	for key, _ := range map1 {
		keySlice = append(keySlice, key)
	}
	fmt.Println(keySlice) //[11 10 4 8 2 1]

	//2、让key进行升序排序
	sort.Ints(keySlice)
	fmt.Println(keySlice) //[1 2 4 8 10 11]

	//3、循环遍历key输出map的值
	for _, v := range keySlice {
		fmt.Printf("key=%v value=%v\n", v, map1[v])
	}

}
