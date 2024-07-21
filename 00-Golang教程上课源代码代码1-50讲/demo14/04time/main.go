package main

import (
	"fmt"
	"time"
)

//时间戳转换成日期字符串
func main() {
	// unixTime: 1587888473

	unixTime := 1587894706
	timeObj := time.Unix(int64(unixTime), 0)
	var str = timeObj.Format("2006-01-02 15:04:05")
	fmt.Println(str) //2020-04-26 17:51:46

}
