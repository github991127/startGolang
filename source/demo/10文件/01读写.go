package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// 读Read
func f1() {
	path := "demo/10文件/test.txt"
	file, err := os.Open(path)
	defer file.Close() // 完成后关闭文件
	if err != nil {
		fmt.Println(err)
		return
	}

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
		strSlice = append(strSlice, tempSlice[:n]...)
	}
	fmt.Println(string(strSlice))
}

// 读bufio.NewReader
func f2() {
	path := "demo/10文件/test.txt"
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	//bufio 读取文件
	var fileStr string
	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') //表示一次读取一行
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

// 写Write
func f3() {
	path := "demo/10文件/writeF3.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println(err)
		return
	}
	// for i := 0; i < 10; i++ {
	// 	file.WriteString("写入String" + strconv.Itoa(i) + "\r\n")
	// }
	var str = "写入byte切片"
	file.Write([]byte(str))

}

// 写bufio.NewWriter
func f4() {
	path := "demo/10文件/writeF4.txt"
	file, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()

	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	writer := bufio.NewWriter(file)

	// writer.WriteString("你好golang") //将数据先写入缓存

	for i := 0; i < 10; i++ {
		writer.WriteString("直接写入的字符串数据" + strconv.Itoa(i) + "\r\n")
	}

	writer.Flush() //将缓存中的内容写入文件
}

func main() {
	//f1()
	f2()
	//f3()
	//f4()
}
