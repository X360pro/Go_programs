package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func main() {
	wg.Add(2)
	ch := make(chan int, 2)
	go WriteToChannel(ch)
	go ReadFromChannel(ch)
	wg.Wait()
}

func WriteToChannel(ch chan int) {

	fmt.Println("Inside GO routine")
	ch <- 10
	ch <- 8
	time.Sleep(time.Second)
	fmt.Println("data has been written to channel")
	wg.Done()
}

func ReadFromChannel(ch chan int) {

	fmt.Println("Starting reading")
	value := <-ch
	value2 := <-ch
	fmt.Println("Data in our channel is : ", value)
	fmt.Println("Data in our channel is : ", value2)
	wg.Done()
}
