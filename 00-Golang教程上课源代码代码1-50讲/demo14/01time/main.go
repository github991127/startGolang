/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"fmt"
	"time"
)

//1、打印当前日期
func main() {

	timeObj := time.Now()
	fmt.Println(timeObj) //2020-04-26 17:32:25.9639049 +0800 CST m=+0.004000301

	// 2020-04-26 17:32:25
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	// fmt.Printf("%d-%d-%d %d:%d:%d", year, month, day, hour, minute, second) //2020-4-26 17:35:07

	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second) //2020-04-26 17:36:04

	//注意：%02d 中的 2 表示宽度，如果整数不够 2 列就补上 0
}
