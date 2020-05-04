package main

import ("fmt"
		"struct")

func main(){
	subscriber := mystruct.Subscriber{Name : "abhishek"}
	subscriber.Home.Street = "MG road"
	subscriber.Home.Pincode = 400045
	fmt.Println(subscriber)
}