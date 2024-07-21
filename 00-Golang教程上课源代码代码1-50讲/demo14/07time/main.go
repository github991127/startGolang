/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"fmt"
	"time"
)

/*
	golang定时器
*/

func main() {

	// // time.Now()
	// ticker := time.NewTicker(time.Second)
	// // ticker.C
	// for t := range ticker.C {
	// 	fmt.Println(t)
	// }

	// time.Now()
	// ticker := time.NewTicker(time.Second)
	// n := 5
	// for t := range ticker.C {
	// 	n--
	// 	fmt.Println(t)
	// 	if n == 0 {
	// 		ticker.Stop() //终止这个定时器继续执行
	// 		break
	// 	}

	// }

	//休眠方法
	// fmt.Println("aaa")
	// time.Sleep(time.Second)
	// fmt.Println("aaa2")
	// time.Sleep(time.Second)
	// fmt.Println("aaa3")
	// time.Sleep(time.Second * 5)
	// fmt.Println("aaa4")

	for {
		time.Sleep(time.Second)
		fmt.Println("我在定时执行任务")
	}
}
