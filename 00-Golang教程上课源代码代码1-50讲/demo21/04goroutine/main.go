package main

import (
	"fmt"
	"sync"
	"time"
)

//多携程的一个案例
var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Printf("协程(%v)打印的第%v条数据\n", num, i)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {
	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Println("关闭主线程...")

}
