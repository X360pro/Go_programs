package main

import "fmt"

func variadic_print(num ... int){
	fmt.Println(num)
}

func main(){
	variadic_print(12,23,4,5,6,78,89)
	println(12,23,4,5,6,78,89)
}