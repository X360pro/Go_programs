package main

import "fmt"

func main() {
	number := 50
	f1()
	defer f2(number)
	number = 100
	defer fmt.Println("End of main")
}

func f1() {
	defer fmt.Println("Hello from f1")
	fmt.Println("end of f1")
}

func f2(number int) {
	fmt.Println("End of f2", number)
}
