package main

import (
	"fmt"
	"sync"
	"time"
)

var wg085 sync.WaitGroup

// var mutex sync.Mutex // 互斥锁
var mutex sync.RWMutex // 读写锁

// 写的方法
func write() {
	mutex.Lock()
	fmt.Println("执行写操作")
	time.Sleep(time.Second * 2)
	mutex.Unlock()
	wg085.Done()
}

// 读的方法
func read() {
	mutex.RLock()
	fmt.Println("---执行读操作")
	time.Sleep(time.Second * 2)
	mutex.RUnlock()
	wg085.Done()
}

func main() {

	//开启10个协程执行读操作
	for i := 0; i < 10; i++ {
		wg085.Add(1)
		go write()
	}
	// 开启10个协程执行写操作
	for i := 0; i < 10; i++ {
		wg085.Add(1)
		go read()
	}
	wg085.Wait()

}
