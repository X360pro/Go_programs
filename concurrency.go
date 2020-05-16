package main

import (
	"fmt"
	"time"
)

func main() {
	go delayedIteration1()
	go delayedIteration2()
	fmt.Scanln()
}

func delayedIteration1() {
	for i := 1; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 1 second : ", i)
	}
}
func delayedIteration2() {
	for i := 1; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("In 2 second : ", i)
	}
}
