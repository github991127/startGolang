package main

import "fmt"

func main() {
	// continue 语句可以结束当前循环，开始下一次的循环迭代过程，仅限在 for 循环内使用。

	// for i := 1; i <= 10; i++ {

	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	// 在 continue 语句后添加标签时，表示开始标签对应的循环

	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 10; j++ {
	// 		if j == 3 {
	// 			continue
	// 		}
	// 		fmt.Printf("i=%v j=%v\n", i, j)
	// 	}
	// }

	/*

	   i=0 j=0
	   i=0 j=1
	   i=0 j=2
	   i=1 j=0
	   i=1 j=1
	   i=1 j=2
	*/

lable2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				continue lable2
			}
			fmt.Printf("i=%v j=%v\n", i, j)
		}
	}
}
