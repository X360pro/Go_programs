package main

import ("time"
		"fmt")

func main(){
	var now time.Time = time.Now()
	fmt.Println("Current time is",now)
	var year int = now.Year()
	fmt.Println("Current year is ",year)
}