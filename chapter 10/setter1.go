package main

import "fmt"

type date struct {
	year  int
	month int
	day   int
}

func (d date) setYearValue(year int) {
	d.year = year
}

func (d *date) setYearPointer(year int) {
	d.year = year
}

func main() {
	newDate := date{}
	newDate.setYearValue(2019)
	fmt.Println("Passing value newDate is ", newDate.year)
	Date := &date{}
	Date.setYearPointer(2019)
	fmt.Println("Passing pointer newDate is ", Date.year)

}
