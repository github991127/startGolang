package main

import "fmt"

// 调用函数
func f1() {
	//func Fn1(x, y int) int {} //形参类型和返回值类型
	//func Fn1(x ...int) int {} //可变参数
	//func Fn1(x int, y ...int) int {} //固定参数和可变参数
	//func Fn1(x, y int) (int, int)  {} //多返回值，return后面跟两个值
	//func Fn1(x, y int) (sum, sub int)  {} //声明变量的多返回值，直接return
}

// 切片作为参数，是引用类型，无须返回也可以直接修改
func f2(slice []int) bool {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] < slice[j] {
				temp := slice[i]
				slice[i] = slice[j]
				slice[j] = temp
			}
		}
	}
	return true

}

// 函数名赋值
type calc func([]int) bool //表示定义一个calc的类型,其是一种函数格式，要求形参和返回类型对应
func f3() {
	var fx calc
	fx = f2
	fmt.Printf("x的类型：%T\n", fx) //x的类型：main.calc
	fy := f2
	fmt.Printf("y的类型：%T\n", fy) //y的类型：func([]int) bool
}

// 函数作为另一个函数参数或返回值
func f4(x, y int, fx calc) calc {
	fmt.Println(fx([]int{x, y})) //调用fx函数
	return fx                    // 返回fx函数，常用于根据输入字符串执行不同的函数

}

// 匿名函数
func f5() {

	// 1.声明并执行
	var fn = func(x, y int) int {
		return x + y
	}
	fmt.Println(fn(2, 3))

	// 2.自执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)
}

func main() {
	//f1() // 调用函数

	//slice := []int{1, 34, 4, 35, 6, 36, 2}
	//f2(slice)
	//fmt.Println(slice)

	//f3()         // 函数名赋值
	//f4(1, 2, f2) // 函数作为另一个函数参数
	f5() // 匿名函数

}
