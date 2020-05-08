package main

import "fmt"

type Mytype string

func (m Mytype) sayHi() {
	fmt.Println("hi ", m)
}

func main() {
	value := Mytype("yo")
	value.sayHi()
}
