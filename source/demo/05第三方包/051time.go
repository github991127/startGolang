package main

import (
	"fmt"
	"time"
)

// time包
func f1() {
	// 获取当前时间
	timeObj := time.Now()
	fmt.Println(timeObj) //2024-04-14 15:18:19.0083427 +0800 CST m=+0.003265901

	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	//%02d 中的 2 表示宽度，如果整数不够 2 列就补上 0
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second) //2024-04-14 15:18:19

	// 时间-时间字符串
	str2 := timeObj.Format("2006-01-02 03:04:05") //12小时制。根据Go的诞生时间2006 1 2 3 4格式化
	fmt.Println(str2)
	str3 := timeObj.Format("2006/01/02 15:04:05") //24小时制
	fmt.Println(str3)

	// 时间字符串-时间
	str4 := "2024-04-14 15:18:19"
	timeObj1, err := time.Parse("2006-01-02 15:04:05", str4)
	if err != nil {
		fmt.Println(timeObj1)
	}

	// 获取Unix时间戳（UnixTimestamp），1970-1-1（08:00:00GMT）至当前时间的总毫秒数
	unixTime := timeObj.Unix() //获取当前的时间戳 （毫秒）
	fmt.Println("当前时间戳：", unixTime)
	unixNatime := timeObj.UnixNano() //纳秒时间戳
	fmt.Println("当前纳秒时间戳：", unixNatime)

	// 时间戳转时间
	timeObj2 := time.Unix(int64(unixTime), 0)
	fmt.Println(timeObj2)

	// 定时器
	ticker := time.NewTicker(time.Second)
	n := 5
	for t := range ticker.C {
		n--
		fmt.Println(t)
		if n == 0 {
			ticker.Stop() //终止这个定时器继续执行
			break
		}
	}
}

func main() {
	f1() // 声明与输出

}
