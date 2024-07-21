package main

import (
	"fmt"
	"time"
)

/*
时间类型有一个自带的方法Format进行格式化，
需要注意的是Go语言中格式化时间模板不是常见的Y-m-d H:M:S
而是使用Go的诞生时间2006年1月2号15点04分（记忆口诀为2006 1 2 3 4）


2006  年
01  月
02  日
03  时   12小时制     写成：15   表示   24小时制
04  分
05  秒

*/
func main() {
	// timeObj := time.Now()
	// var str = timeObj.Format("2006-01-02 03:04:05")
	// fmt.Println(str) //2020-04-26 05:44:24

	//12小时制
	// timeObj := time.Now()
	// var str = timeObj.Format("2006/01/02 03:04:05")
	// fmt.Println(str) //2020/04/26 05:44:47

	// 24小时制
	timeObj := time.Now()
	var str = timeObj.Format("2006/01/02 15:04:05")
	fmt.Println(str) // 2020/04/26 17:48:53
}
