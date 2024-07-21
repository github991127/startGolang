package main

import "fmt"

func main() {
	//如果我们想在map对象中存放一系列的属性的时候，我们就可以把map类型的值定义成切片

	// var userinfo = make(map[string]string)
	// userinfo["username"] = "张三"
	// userinfo["hobby"] = "睡觉"

	var userinfo = make(map[string][]string)

	userinfo["hobby"] = []string{
		"吃饭",
		"睡觉",
		"敲代码",
	}

	userinfo["work"] = []string{
		"php",
		"golang",
		"前端",
	}

	// fmt.Println(userinfo)

	for _, v := range userinfo {
		// fmt.Println(k, v)
		for _, value := range v {
			fmt.Println(value)
		}
	}
}
