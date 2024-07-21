package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 1、创建channel
	var ch1 = make(chan int, 3)
	wg.Add(1)
	go func() {
		for i := 1; i <= 3; i++ {
			num := <-ch1
			fmt.Println(num)
		}
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		for i := 1; i <= 3; i++ {
			time.Sleep(time.Second)
			ch1 <- i
		}
		wg.Done()
	}()

	wg.Wait()
}
