/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"errors"
	"fmt"
)
/*
panic/recover

Go 语言中目前（Go1.12）是没有异常机制，但是使用 panic/recover 模式来处理错误。

panic 可以在任何地方引发，但 recover 只有在 defer 调用的函数中有效
*/

// 模拟了一个读取文件的方法
func readFile(fileName string) error {
	if fileName == "main.go" {
		return nil
	} else {
		return errors.New("读取文件失败")
	}
}

func myFn() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("给管理员发送邮件")
		}
	}()
	err := readFile("main.go")
	if err != nil {
		panic(err)
	}
}

func main() {
	myFn()
	fmt.Println("继续执行...")
}
