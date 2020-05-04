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
	person1.firstname = "abhishek"
	person1.lastname = "nagvekar"
	person1.age = 20
	person1.address = make([]string,2)
	person1.address[0] = "andheri"
	fmt.Println(person1)
}