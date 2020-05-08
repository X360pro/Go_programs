package main

import "fmt"

type Liters float32
type Gallons float32

func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(2.5)
	busFuel = ToLiters(Gallons(8.0))
	fmt.Println(carFuel, busFuel)
}
