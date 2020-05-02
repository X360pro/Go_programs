package main

import "fmt"

func arr_div(array []int) []float64 {
	limit := len(array)
	newArray := make([]float64,limit)
	for i,number := range array {
		newArray[i] = float64(number/3)	
	}
	return newArray
}

func main(){
	array := []int{3,6,9,12,15}
	answer := arr_div(array)
	fmt.Println(answer)

}