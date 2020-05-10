package main

import "fmt"

type coordinates struct {
	lat float64
	lon float64
}

func (c *coordinates) setLat(latitude float64) {
	c.lat = latitude
}

func (c *coordinates) setLon(longitude float64) {
	c.lon = longitude
}

func main() {
	newCoordinates := &coordinates{}
	newCoordinates.setLat(23)
	newCoordinates.setLon(55)
	fmt.Println(newCoordinates)
}
