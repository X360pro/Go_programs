package main

import "fmt"

type subscriber struct{
	name string
	rate float64
	active bool
}

func main(){
	var subscriber1 subscriber
	var subscriber2 subscriber
	subscriber1.name = "Abhishek"
	subscriber2.name = "Soham"
	subscriber1.rate = 1234
	subscriber2.rate = 2345
	subscriber1.active = false
	subscriber2.active = true
	fmt.Println(subscriber1,subscriber2)
}