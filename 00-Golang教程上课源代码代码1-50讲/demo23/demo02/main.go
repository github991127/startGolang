package main

import (
	"fmt"
	"io"
	"os"
)

/*
	一、读取文件（方法1）
		1、只读方式打开文件 file,err := os.Open()

		2、读取文件 file.Read()

		3、关闭文件流 defer file.Close()
*/

func main() {
	//1、打开文件
	file, err := os.Open("C:/test.txt")
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	//2、读取文件里面的内容
	var strSlice []byte
	var tempSlice = make([]byte, 128)
	for {
		n, err := file.Read(tempSlice)
		if err == io.EOF { //err==io.EOF表示读取完毕
			fmt.Println("读取完毕")
			break
		}
		if err != nil {
			fmt.Println("读取失败")
			return
		}
		// fmt.Printf("读取到了%v个字节\n", n)
		strSlice = append(strSlice, tempSlice[:n]...) //注意写法
	}

	fmt.Println(string(strSlice))
}
