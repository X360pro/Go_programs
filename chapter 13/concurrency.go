package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println(runtime.GOMAXPROCS(-1))
	runtime.GOMAXPROCS(2)
	fmt.Println(runtime.GOMAXPROCS(-1))
	wg.Add(2)
	go delayedIteration1()
	go delayedIteration2()
	wg.Wait()
}

func delayedIteration1() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 1 second : ", i)
	}
	wg.Done()
}
func delayedIteration2() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 2 second : ", i)
	}
	wg.Done()
}
