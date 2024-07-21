package main

import (
	"fmt"
	"time"
)
//需求：要统计1-120000的数字中那些是素数？for循环实现
func main() {

	start := time.Now().Unix()
	for num := 2; num < 120000; num++ {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(num, "是素数")
		}
	}
	end := time.Now().Unix()

	fmt.Println(end - start) //11毫秒  11毫秒

}
