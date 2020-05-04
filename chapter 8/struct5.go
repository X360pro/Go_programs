package main

import "fmt"

type coordinates struct {
	lat float64
	lon float64
}

func main(){
	location := coordinates{lat:37.42,lon:-122.08}
	fmt.Println(location)
}