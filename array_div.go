package main

import "fmt"

func arr_div(array [5]int) [5]float64 {
	var newArray [5]float64
	for i,number := range array {
		newArray[i] = float64(number/3)	
	}
	return newArray
}

func main(){
	array := [5]int{3,6,9,12,15}
	answer := arr_div(array)
	fmt.Println(answer)
}
