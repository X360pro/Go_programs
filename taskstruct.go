package main

import "fmt"

type person struct {
	firstname string
	lastname string
	age int8
	address []string
}

func main(){
	var person1 person
	var person2 person
	person1.firstname = "abhishek"
	person1.lastname = "nagvekar"
	person1.age = 20
	person1.address = make([]string,1)
	person1.address[0] = "andheri"
	fmt.Println(person1)
	person2.firstname = "soham"
	person2.lastname = "nagvekar"
	person2.age = 16
	person2.address = make([]string,1)
	person2.address[0] = "andheri"
	fmt.Println(person2)

	person3 := person{firstname:"Max",lastname:"John",age:40,address : []string{"New York"}}
	fmt.Println(person3)
}
