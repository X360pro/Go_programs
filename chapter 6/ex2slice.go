package main

import ("fmt"
		"os")

func main(){
	limit := os.Args[1:]
	fmt.Println(limit)
	array := [5]string{"a","b","c","d","e"}
	slice := array[1:3]
	slice = append(slice,"x")
	slice = append(slice,"y","z")
	for _,letter := range slice {
		fmt.Println(letter)
	}
}