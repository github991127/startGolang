package main

import (
	"fmt"
	"sync"
	"time"
)

var wg083 sync.WaitGroup

func putNum(intChan chan int) {
	for i := 2; i < 1000; i++ {
		intChan <- i
	}
	close(intChan)
	wg083.Done()
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	for num := range intChan {
		var flag = true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag {
			primeChan <- num
		}
	}
	//需要在所有协程都退出后关闭 primeChan
	//close(primeChan)//后续打印协程还需要读取primeChan，所以不能关闭
	exitChan <- true //通过标识记录当前已完成协程数，若所有协程都退出，则exitChan会存满值，关闭primeChan
	wg083.Done()
}

func printPrime(primeChan chan int) {
	for v := range primeChan {
		fmt.Println(v)
	}
	wg083.Done()
}

func main() {

	start := time.Now().Unix()
	primeCount()
	end := time.Now().Unix()
	fmt.Println("执行完毕....", end-start, "毫秒")

}

func primeCount() {
	coroutineNum := 4
	intChan := make(chan int, 1000)           //存放所有数字
	primeChan := make(chan int, 1000)         //存放素数
	exitChan := make(chan bool, coroutineNum) //标识primeChan close

	//存放数字的协程
	wg083.Add(1)
	go putNum(intChan)

	//统计素数的协程
	for i := 0; i < coroutineNum; i++ {
		wg083.Add(1)
		go primeNum(intChan, primeChan, exitChan) //intChan一边写，一边读入primeChan
	}

	//打印素数的协程
	wg083.Add(1)
	go printPrime(primeChan)

	//判断exitChan是否存满值
	wg083.Add(1)
	go func() {
		for i := 0; i < coroutineNum; i++ {
			<-exitChan
		}
		close(primeChan)
		wg083.Done()
	}()

	wg083.Wait()
}
