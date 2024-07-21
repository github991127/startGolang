package main

import (
	"fmt"
	"sort"
	"strings"
)

// 字典
func f5() {
	//初始化
	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["age"] = "20"
	// userinfo["sex"] = "男"
	// // fmt.Println(userinfo)
	// fmt.Println(userinfo["username"])
	var userinfo = map[string]string{
		"username": "张三",
		"age":      "20",
		"sex":      "男",
	}
	fmt.Println(userinfo)
	for k, v := range userinfo {
		fmt.Printf("key:%v value:%v\n", k, v)
	}

	//修改，map类型是引用数据类型
	userinfo["age"] = "21"
	//删除
	delete(userinfo, "sex")
	//查找
	v, b := userinfo["age"]
	fmt.Println(v, b) //21 true
	v2, b2 := userinfo["email"]
	fmt.Println(v2, b2) //"" false

}

// 字典切片
func f6() {
	//如果我们想存放多组map对象，我们就可以使用切片来存放
	var userinfo = make([]map[string]string, 2, 2)
	fmt.Println(userinfo[0])        //map[]   map不初始化的默认值nil
	fmt.Println(userinfo[0] == nil) //true

	if userinfo[0] == nil {
		userinfo[0] = make(map[string]string)
		userinfo[0]["username"] = "张三"
		userinfo[0]["age"] = "20"
		userinfo[0]["height"] = "180cm"
	}

	if userinfo[1] == nil {
		userinfo[1] = make(map[string]string)
		userinfo[1]["username"] = "李四"
		userinfo[1]["age"] = "22"
		userinfo[1]["height"] = "170cm"
	}
	fmt.Println(userinfo)

	//如果我们想在map对象中存放一系列的属性的时候，我们就可以把map类型的值定义成切片
	var userinfo2 = make(map[string][]string)
	userinfo2["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}

	userinfo2["work"] = []string{
		"php",
		"golang",
		"前端",
	}
	fmt.Println(userinfo2)

}

// 字典排序
func f7() {
	//key排序并输出key-value
	nums := map[int]int{
		int(11): 100,
		int(10): 200,
		int(4):  300,
		int(8):  400,
		int(2):  500,
	}
	//1、把map的key放在切片里面
	var keySlice []int
	for key, _ := range nums {
		keySlice = append(keySlice, key)
	}
	fmt.Println(keySlice)

	//2、让key进行升序排序
	sort.Ints(keySlice)
	fmt.Println(keySlice)

	//3、循环遍历key输出map的值
	for _, v := range keySlice {
		fmt.Printf("key=%v value=%v\n", v, nums[v])
	}
}

// 统计单词数
func f8() {
	words := "go go go go is a good language"
	strSlice := strings.Split(words, " ")
	wordCount := make(map[string]int)

	for _, v := range strSlice {
		wordCount[v]++
	}

	fmt.Println(wordCount)

}
func main() {
	//f5() // 字典
	//f6() // 字典切片
	//f7() // 字典排序
	f8() // 统计单词数
}
