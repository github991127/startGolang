package main

import (
	"fmt"
	"os"
)
func main() {

	// err := os.Mkdir("./abc", 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// os.Mkdir("./abc", 0666)

	err := os.MkdirAll("./dir1/dir2/dir3", 0666) //创建多级目录
	if err != nil {
		fmt.Println(err)
	}
}
