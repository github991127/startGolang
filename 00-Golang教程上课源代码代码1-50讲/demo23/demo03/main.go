package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
二、读取文件（方法2）bufio 读取文件

		1、只读方式打开文件 file,err := os.Open()

		2、创建reader对象  reader := bufio.NewReader(file)

		3、ReadString读取文件  line, err := reader.ReadString('\n')

		4、关闭文件流 defer file.Close()
*/

func main() {
	file, err := os.Open("C:/test.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	//bufio 读取文件
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //表示一次读取以行
		if err == io.EOF {
			fileStr += str
			break
		}
		if err != nil {
			fmt.Println(err)
			return
		}
		fileStr += str
	}
	fmt.Println(fileStr)

}
