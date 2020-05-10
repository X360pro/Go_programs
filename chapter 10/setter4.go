package main

import (
	"fmt"
	"geo"
	"log"
)

func main() {
	location := geo.Landmark{}
	err := location.SetName("The Google plex")
	if err != nil {
		log.Fatal(err)
	}
	location.SetLat(37.42)
	location.SetLon(-122.08)
	fmt.Println(location)
}
