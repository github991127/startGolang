package main

import (
	"fmt"
	"time"
)

//日期字符串转换成时间戳
func main() {

	var str = "2020-04-26 15:38:04"
	var tmp = "2006-01-02 15:04:05"
	timeObj, _ := time.ParseInLocation(tmp, str, time.Local)

	fmt.Println(timeObj)        //2020-04-26 15:38:04 +0800 CST
	fmt.Println(timeObj.Unix()) //1587886684

}
