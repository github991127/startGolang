package main

import (
	"fmt"
	"io/ioutil"
)

// ioutil按行读取示例
func main() {
	byteStr, err := ioutil.ReadFile("C:/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(byteStr))
}
