package main

import (
	"github.com/headfirstgo/gadget"
)

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

func main() {
	player := gadget.TapePlayer{}
	mixtape := []string{"Jessies Girl", "Whip it", "9 to 5"}
	playList(player, mixtape)
}
