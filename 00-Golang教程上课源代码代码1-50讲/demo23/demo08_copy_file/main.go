package main

import (
	"fmt"
	"io/ioutil"
)

//自己编写一个函数，接收两个文件路径 srcFileName dstFileName

func copy(srcFileName string, dstFileName string) (err error) {
	byteStr, err1 := ioutil.ReadFile(srcFileName)
	if err1 != nil {
		return err1
	}
	err2 := ioutil.WriteFile(dstFileName, byteStr, 0666)
	if err2 != nil {

		return err2
	}
	return nil
}

func main() {
	src := "C:/test.txt"
	dst := "D:/test1.txt"
	err := copy(src, dst)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("复制文件成功")
}
