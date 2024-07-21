package main

//go build -race main.go  编译后运行查看

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup

var mutex sync.Mutex

func test() {
	mutex.Lock()
	count++
	fmt.Println("the count is : ", count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()

}
