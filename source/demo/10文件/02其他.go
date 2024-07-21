package main

import (
	"fmt"
	"os"
)

// 重命名文件
func f5() {
	path := "demo/10文件/testF5.txt"
	pathNew := "demo/10文件/testF555.txt"
	err := os.Rename(path, pathNew)
	if err != nil {
		fmt.Println(err)
	}
}

// 创建多级目录
func f6() {
	err := os.MkdirAll("./dir1/dir2/dir3", 0666) //创建多级目录
	if err != nil {
		fmt.Println(err)
	}
}

// 删除
func f7() {
	//删除一个
	err := os.Remove("aaa.txt")
	if err != nil {
		fmt.Println(err)
	}

	//删除多个
	//err := os.RemoveAll("dir1")
	//if err != nil {
	//	fmt.Println(err)

}
func main() {
	//f5()
	//f6()
	//f7()

}
