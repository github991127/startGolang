/*
	大地老师专栏：https://www.itying.com/category-79-b0.html
	Golang仿小米商城项目实战视频教程地址：https://www.itying.com/goods-1143.html
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mutex sync.RWMutex

//写的方法
func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg.Done()
}

//读的方法
func read() {
	mutex.RLock()
	fmt.Println("---执行读操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg.Done()
}

func main() {

	//开启10个协程执行读操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	// 开启10个协程执行写操作
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()

}
