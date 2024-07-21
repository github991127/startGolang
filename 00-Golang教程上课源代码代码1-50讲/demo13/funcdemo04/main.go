package main

import (
	"fmt"
	"sort"
)

func mapSort(map1 map[string]string) string {

	var sliceKey []string

	//1、把map对象的key放在一个切片里面
	for k, _ := range map1 {
		sliceKey = append(sliceKey, k)
	}
	//2、对key进行升序排序
	sort.Strings(sliceKey)

	var str string
	for _, v := range sliceKey {
		str += fmt.Sprintf("%v=>%v", v, map1[v])
	}
	return str

}

func main() {
	/*
		案例2：
				m1 := map[string]string{
					"username":"zhangsan",
					"age":"20",
					"sex":"男",
					"height":"180",
				}
				输出：age=>20height=>180sex=>男username=>zhangsan

				上面有一个map对象m1,封装一个方法,要求按照key升序排序，最后输出一个key=>valuekey=>value的字符串


	*/

	m1 := map[string]string{
		"username": "zhangsan",
		"age":      "20",
		"sex":      "男",
		"height":   "180",
		"salt":     "xxxxxx",
	}

	str := mapSort(m1)
	fmt.Println(str)
}
