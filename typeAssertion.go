package main

import "fmt"

func main() {
	var i interface{} = 10
	storeString, ok := i.(int)
	fmt.Println(storeString, ok)
	checkType(true)
}

func checkType(anyType interface{}) {
	switch anyType.(type) {
	case string:
		fmt.Println("arguement is a string")

	case float32:
		fmt.Println("arguement is a float")

	case int:
		fmt.Println("arguement is an integer")

	default:
		fmt.Println("default is chosen")
	}
}
