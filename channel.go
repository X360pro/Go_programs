package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go WriteTChannel(ch)
	value := <-ch
	fmt.Println("Data in our channel is : ", value)
}
func WriteTChannel(ch chan int) {

	fmt.Println("Inside GO routine")
	ch <- 10
	time.Sleep(time.Second)
	fmt.Println("data has been written to channel")
}
