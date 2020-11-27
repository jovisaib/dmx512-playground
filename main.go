package main

import (
	"log"

	"github.com/akualab/dmx"
)

func main() {

	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal(e)
	}


	for {
		// Set values for channels.
		dmx.SetChannel(1, 100)
		dmx.SetChannel(2, 70)
		dmx.SetChannel(3, 130)
		dmx.SetChannel(4, 180)
		
		dmx.SetChannel(5, 100)
		dmx.SetChannel(6, 70)
		dmx.SetChannel(7, 130)
		dmx.SetChannel(8, 180)
	
		dmx.SetChannel(9, 100)
		dmx.SetChannel(10, 70)
		dmx.SetChannel(11, 130)
		dmx.SetChannel(12, 180)
	
		dmx.SetChannel(13, 100)
		dmx.SetChannel(14, 70)
		dmx.SetChannel(15, 130)
		dmx.SetChannel(16, 180)
	
		dmx.SetChannel(26, 100)
		dmx.SetChannel(27, 70)
		dmx.SetChannel(28, 130)
		dmx.SetChannel(29, 180)
	
		dmx.SetChannel(30, 100)
		dmx.SetChannel(31, 70)
		dmx.SetChannel(32, 130)
		dmx.SetChannel(33, 180)
	
		dmx.SetChannel(34, 100)
		dmx.SetChannel(35, 70)
		dmx.SetChannel(36, 130)
		dmx.SetChannel(37, 180)
	
		dmx.SetChannel(38, 100)
		dmx.SetChannel(39, 70)
		dmx.SetChannel(40, 130)
		dmx.SetChannel(41, 180)
	
		// Send!
		dmx.Render()
	}
}
