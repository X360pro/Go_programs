package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	age       int8
	address   []string
}

func assign_values(curPerson person, first, last string, curAge int8, curAdd []string) person {
	curPerson.firstname = first
	curPerson.lastname = last
	curPerson.age = curAge
	curPerson.address = curAdd
	return curPerson
}

func display(curPerson person) {
	fmt.Println("First Name : ", curPerson.firstname)
	fmt.Println("Last Name : ", curPerson.lastname)
	fmt.Println("Age : ", curPerson.age)
	fmt.Println(curPerson.firstname, " has address in ")
	for _, curAddress := range curPerson.address {
		fmt.Println("        ", curAddress)
	}
}

func main() {
	var person1 person
	addresses := []string{"usa", "india", "uk", "canada"}
	person1 = assign_values(person1, "john", "lark", 46, addresses)
	display(person1)
}
