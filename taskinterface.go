package main

import "fmt"

type shape interface {
	perimeter() float64
	area() float64
}

type square struct {
	side float64
}

func (s *square) perimeter() float64 {
	return (4 * s.side)
}

func (s *square) area() float64 {
	return (s.side * s.side)
}

func printDetals(s shape) {
	fmt.Println("Perimeter of square is ", s.perimeter())
	fmt.Println("Area is ", s.area())

}

func main() {
	var newSquare shape
	newSquare = &square{5.0}
	printDetals(newSquare)
}
