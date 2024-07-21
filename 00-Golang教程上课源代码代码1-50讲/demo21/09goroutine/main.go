package main

import (
	"fmt"
	"sync"
)
package main
var wg sync.WaitGroup
func putNum(intChan chan int) {
	for i := 2; i < 10; i++ {
		intChan <- i
		fmt.Println(i)
	}
	close(intChan)
	wg.Done()
}
//定义一个存放任意数据类型的管道 3个数据
func main() {

	// allChan := make(chan interface{}, 3)

	// allChan <- 10
	// allChan <- "tom jack"
	// cat := Cat{"小花猫", 4}
	// allChan <- cat

	// //我们希望获得到管道中的第三个元素，则先将前2个推出
	// <-allChan
	// <-allChan

	// newCat := <-allChan //从管道中取出的Cat是什么?

	// fmt.Printf("newCat=%T , newCat=%v\n", newCat, newCat)
	// //下面的写法是错误的!编译不通过
	// //fmt.Printf("newCat.Name=%v", newCat.Name)
	// //使用类型断言
	// a := newCat.(Cat)
	// fmt.Printf("newCat.Name=%v", a.Name)

	var intChan = make(chan int, 1000)
	wg.Add(1)
	go putNum()
	wg.Wait()()
}
