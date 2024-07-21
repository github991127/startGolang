package main

import (
	"fmt"
	"os"
)

/*
	一、写入文件（方法1）
		1、打开文件  file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_RDWR, 0666)

		2、写入文件
			file.Write([]byte(str))        //写入字节切片数据
			file.WriteString("直接写入的字符串数据") //直接写入字符串数据

		3、关闭文件流 file.Close()
*/
func main() {
	file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	//写如文件
	// for i := 0; i < 10; i++ {
	// 	file.WriteString("直接写入的字符串数据" + strconv.Itoa(i) + "\r\n")
	// }
	var str = "直接写入的字符串数据byte"
	file.Write([]byte(str))
}
