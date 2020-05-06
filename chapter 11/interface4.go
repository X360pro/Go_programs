package main

import "fmt"

type Switch string

func (s *Switch) toggle() {
	if *s == "on" {
		*s = "off"
	} else {
		*s = "on"
	}
	fmt.Println(*s)
}

type toggleable interface {
	toggle()
}

func main() {
	s := Switch("off")
	var t toggleable = &s
	t.toggle()
	t.toggle()
}
