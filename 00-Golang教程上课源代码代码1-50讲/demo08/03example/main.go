package main

import "fmt"

func main() {
	
	//1、练习：打印 0-50 所有的偶数

	// for i := 0; i <= 50; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }

	//2、练习：求 1+2+3+4 +...100 的和
	/*
		sum += sum+1   0+1

		sum += sum+2   0+1+2

		sum += sum+3   0+1+2+3

		sum += sum+4   0+1+2+3+4
	*/

	// sum := 0
	// for i := 1; i <= 100; i++ {
	// 	sum += i
	// }
	// fmt.Println(sum) //5050

	//3、练习：打印 1~100 之间所有是 9 的倍数的整数的个数及总和

	// var sum = 0
	// var count = 0
	// for i := 1; i <= 100; i++ {
	// 	if i%9 == 0 {
	// 		fmt.Println(i)
	// 		sum += i
	// 		count++
	// 	}
	// }

	// fmt.Println(sum)
	// fmt.Println(count)

	//4、练习：计算 5 的阶乘 (12345 n 的阶乘 12……n)   1*2*3*4*5

	/*
		1、sum=1*1
		2、sum=sum*2   1*1*2
		3、sum=sum*3   1*1*2*3
		4、sum=sum*4   1*1*2*3*4
		5、sum=sum*5   1*1*2*3*4*5
	*/

	// var sum = 1
	// for i := 1; i <= 5; i++ {
	// 	sum *= i
	// }
	// fmt.Println(sum)

	//5、练习： 打印一个矩形 (for循环的嵌套)

	/*

	****
	****
	****

	 */

	// for i := 1; i <= 12; i++ {
	// 	fmt.Print("*")
	// 	if i%4 == 0 {
	// 		fmt.Println("")
	// 	}
	// }

	//6、for循环的嵌套

	/*
		for循环嵌套的一个执行流程
		1、i=0   打印4个*   一个换行
		2、i=1   打印4个*   一个换行
		3、i=2   打印4个*   一个换行
		4、i=3  跳出循环
	*/

	// var row = 3
	// var column = 4
	// for i := 0; i < row; i++ {
	// 	for j := 0; j < column; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	//7、练习： 打印一个三角形
	/*

	*
	**
	***
	****
	*****

	 */

	// var row = 5
	// for i := 1; i <= row; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println("")
	// }

	// 8、练习：打印出九九乘法表

	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v*%v=%v \t", j, i, i*j)
		}
		fmt.Println("")
	}

}
