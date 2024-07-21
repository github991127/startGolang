package main

//go build -race main.go  编译后运行查看

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mutex sync.Mutex

var m = make(map[int]int, 0)

func test(num int) {
	mutex.Lock()
	var sum = 1
	for i := 1; i <= num; i++ {
		sum *= i
	}
	m[num] = sum
	// fmt.Println(m[num])
	fmt.Printf("key=%v value=%v\n", num, sum)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for r := 0; r < 40; r++ {
		wg.Add(1)
		go test(r)
	}
	wg.Wait()

}
