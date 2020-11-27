package main

import (
	"log"

	"github.com/akualab/dmx"
)

func main() {

	dmx, e := dmx.NewDMXConnection("/dev/tty3")
	if e != nil {
		log.Fatal(e)
	}

	// Set values for channels.
	dmx.SetChannel(1, 100)
	dmx.SetChannel(2, 70)
	dmx.SetChannel(3, 130)
	dmx.SetChannel(4, 180)

	// Send!
	dmx.Render()
}
