package main

import (
	"fmt"
	"strings"
)

func main() {
	// 写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。

	// this is golang

	var str = "how do you do"
	var strSlice = strings.Split(str, " ")
	fmt.Println(strSlice)

	var strMap = make(map[string]int)
	for _, v := range strSlice {
		strMap[v]++
	}
	fmt.Println(strMap)

}
