package main

import "fmt"

/*
	Go 语言的内建函数 append()可以为切片动态添加元素，每个切片会指向一个底层数组，这 个数组的容量够用就添加新增元素。当底层数组不能容纳新增的元素时，切片就会自动按照 一定的策略进行“扩容”，此时该切片指向的底层数组就会更换。“扩容”操作往往发生在 append()函数调用时，所以我们通常都需要用原变量接收 append 函数的返回值
*/

func main() {

	//1、append方法的使用
	// var sliceA []int

	// sliceA = append(sliceA, 12)
	// sliceA = append(sliceA, 24)
	// fmt.Printf("%v - %v--%v", sliceA, len(sliceA), cap(sliceA)) //[12 24] - 2--2

	// var sliceA []int
	// sliceA = append(sliceA, 12, 23, 35, 465)
	// fmt.Printf("%v - %v--%v", sliceA, len(sliceA), cap(sliceA)) //[12 23 35 465] - 4--4

	//2、append方法还可以合并切片

	// sliceA := []string{"php", "java"}
	// sliceB := []string{"nodejs", "go"}
	// sliceA = append(sliceA, sliceB...)
	// fmt.Println(sliceA) //[php java nodejs go]

	//3、切片的扩容策略(了解)
	/*
		1、首先判断，如果新申请容量（cap）大于 2 倍的旧容量（old.cap），最终容量（newcap） 就是新申请的容量（cap）。
		2、否则判断，如果旧切片的长度小于 1024，则最终容量(newcap)就是旧容量(old.cap)的两 倍，即（newcap=doublecap），
		3、否则判断，如果旧切片长度大于等于 1024，则最终容量（newcap）从旧容量（old.cap） 开始循环增加原来的 1/4，即（newcap=old.cap,for {newcap += newcap/4}）直到最终容量 （newcap）大于等于新申请的容量(cap)，即（newcap >= cap）
		4、如果最终容量（cap）计算值溢出，则最终容量（cap）就是新申请容量（cap）。

		需要注意的是，切片扩容还会根据切片中元素的类型不同而做不同的处理，比如 int 和 string 类型的处理方式就不一样。
	*/
	var sliceA []int
	for i := 1; i <= 10; i++ {
		sliceA = append(sliceA, i)
		fmt.Printf("%v 长度:%d 容量:%d\n", sliceA, len(sliceA), cap(sliceA))
	}
}
