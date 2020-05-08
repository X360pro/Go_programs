package main

import "fmt"

type Mytype string

func (m Mytype) sayHi() {
	fmt.Println("hi ", m)
}

func (m Mytype) Parameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}
func (m Mytype) Return() int {
	return len(m)
}
func main() {
	value := Mytype("yo")
	value.sayHi()
	value.Parameters(4, true)
	fmt.Println(value.Return())
}
