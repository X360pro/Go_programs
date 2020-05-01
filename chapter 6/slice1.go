package main

import ("fmt")

func main(){
	var mySlice []string
	mySlice = make([]string,7)
	mySlice[0] = "I"
	mySlice[1] = "am"
	mySlice[2] = "abhishek"
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	for _,letter := range mySlice {
		fmt.Println(letter)
	}
}