package main

import (
	"fmt"
	"time"
)

//日期字符串转换成时间戳
func main() {

	/*
		1、time包中定义的时间间隔类型的常量如下：
			const (
			    Nanosecond  Duration = 1
			    Microsecond          = 1000 * Nanosecond
			    Millisecond          = 1000 * Microsecond
			    Second               = 1000 * Millisecond
			    Minute               = 60 * Second
			    Hour                 = 60 * Minute
			)
	*/
	// fmt.Println(time.Millisecond) //1毫秒
	// fmt.Println(time.Second)      //1秒

	/*
		2、时间操作函数
	*/

	var timeObj = time.Now()
	fmt.Println(timeObj)
	timeObj = timeObj.Add(time.Hour)
	fmt.Println(timeObj)

	/*
		2020-04-26 18:15:05.4612997 +0800 CST m=+0.005000201
		2020-04-26 19:15:05.4612997 +0800 CST m=+3600.005000201
	*/

}
