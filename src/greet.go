package main

import ("time"
		"fmt"
		"strconv")
func greet(now time.Time) string {
	var currentTimestr = now.Format("15")
	var currentTime,_ = strconv.Atoi(currentTimestr)
	fmt.Println(currentTime)
	if currentTime >=0 && currentTime < 6  {
		return "Good Night"
	} else if currentTime >= 6 && currentTime < 12 {
		return "Good Morning"
	} else if currentTime >= 12 && currentTime < 17{
		return "Good Afternoon" 
	} else if currentTime >= 17 &&  currentTime < 19 {
		return "Good Evening"
	} else if currentTime >=19 && currentTime <= 23 {
		return "Good Night"
	}
	return "working"
}
package main

import ("time"
		"fmt"
		"strconv")
func greet(now time.Time) string {
	var currentTimestr = now.Format("15")
	var currentTime,_ = strconv.Atoi(currentTimestr)
	fmt.Println(currentTime)
	if currentTime >=0 && currentTime < 6  {
		return "Good Night"
	} else if currentTime >= 6 && currentTime < 12 {
		return "Good Morning"
	} else if currentTime >= 12 && currentTime < 17{
		return "Good Afternoon"
	} else if currentTime >= 17 &&  currentTime < 19 {
		return "Good Evening"
	} else if currentTime >=19 && currentTime <= 23 {
		return "Good Night"
	}
	return "working"
}
func main(){
	var now time.Time = time.Now()
	fmt.Println("Current time is",now)
	greeting := greet(now)
	fmt.Println(greeting)
	
}func main(){
	var now time.Time = time.Now()
	fmt.Println("Current time is",now)
	greeting := greet(now)
	fmt.Println(greeting)
	
}