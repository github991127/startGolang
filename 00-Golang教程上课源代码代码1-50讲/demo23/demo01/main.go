/*
读取文件：
	一、读取文件（方法1）
		1、只读方式打开文件 file,err := os.Open()

		2、读取文件 file.Read()

		3、关闭文件流 defer file.Close()

	二、读取文件（方法2）bufio 读取文件

		1、只读方式打开文件 file,err := os.Open()

		2、创建reader对象  reader := bufio.NewReader(file)

		3、ReadString读取文件  line, err := reader.ReadString('\n')

		4、关闭文件流 defer file.Close()

	三、读取文件（方法3）ioutil 读取文件

		打开关闭文件的方法它都封装好了只需要一句话就可以读取

		ioutil.ReadFile("./main.go")


写入文件：

	一、写入文件（方法1）
		1、打开文件  file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_RDWR, 0666)

		2、写入文件
			file.Write([]byte(str))        //写入字节切片数据
			file.WriteString("直接写入的字符串数据") //直接写入字符串数据

		3、关闭文件流 file.Close()

	二、写入文件（方法2） bufio 写入文件

		1、打开文件  file, err := os.OpenFile("C:/test.txt", os.O_CREATE|os.O_RDWR, 0666)

		2、创建writer对象  writer := bufio.NewWriter(file)

		3、将数据先写入缓存  writer.WriteString("你好golang\r\n")

		4、将缓存中的内容写入文件	writer.Flush()

		5、关闭文件流 file.Close()

	三、写入文件（方法3） ioutil 写入文件

		str := "hello golang"
		err := ioutil.WriteFile("C:/test.txt", []byte(str), 0666)

文件重命名：

	err := os.Rename("C:/test1.txt", "D:/test1.txt") //只能同盘操作

复制文件：
	一、方法1
		input, err := ioutil.ReadFile(srcFileName)

		err = ioutil.WriteFile(dstFileName, input, 0644)

	二、方法2

			source, _ := os.Open(srcFileName)

			destination, _ := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)

			n, err := source.Read(buf)

			destination.Write(buf[:n]);
创建目录

		err := os.Mkdir("./abc", 0666)

		err := os.MkdirAll("dir1/dir2/dir3", 0666) //创建多级目录


删除目录和文件

		err := os.Remove("t.txt")

		err := os.RemoveAll("aaa")
*/

package main

import (
	"fmt"
	"os"
)

func main() {
	// 只读方式打开当前目录下的main.go文件  D:/go_demo/demo23/demo01/main.go
	file, err := os.Open("./main.go")
	defer file.Close() //必须得关闭文件流
	if err != nil {
		fmt.Println(err)
		return
	}
	//操作文件
	fmt.Println(file) //&{0xc00014a780}
}
