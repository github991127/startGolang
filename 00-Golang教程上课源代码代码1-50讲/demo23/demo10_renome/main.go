package main

import (
	"fmt"
	"os"
)
func main() {
	err := os.Rename("D:/000.avi", "D:/1111.avi")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("重命名成功")
}
