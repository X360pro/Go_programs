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
	ch2 := make(chan struct{})
	go WriteToChannel(ch)
	go ReadFromChannel(ch)
	wg.Wait()
	close(ch)
	go func() {
		ch2 <- struct{}{}
	}()
	for i := range ch {
		fmt.Println("Inside range", i)
	}

	for i := 0; i < 2; i++ {
		select {
		case chan1, ok := <-ch:
			fmt.Println("channel 1 in data : ", chan1, ok)
		case chan2, ok := <-ch2:
			fmt.Println("channel 2 in data : ", chan2, ok)
		}
	}
}

func WriteToChannel(ch chan int) {

	fmt.Println("Inside GO routine")
	ch <- 10
	ch <- 8
	ch <- 6
	time.Sleep(time.Second)
	fmt.Println("data has been written to channel")
	wg.Done()
}

func ReadFromChannel(ch chan int) {

	fmt.Println("Starting reading")
	value, ok := <-ch
	value2 := <-ch
	fmt.Println("Data in our channel is : ", value, ok)
	fmt.Println("Data in our channel is : ", value2)
	wg.Done()
}
