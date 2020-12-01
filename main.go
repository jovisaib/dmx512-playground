package main

import (
	"log"
	"time"

	"github.com/jovisaib/dmx512-playground/dmx"
)

func main() {

	// r := rand.New(rand.NewSource(99))
	log.Printf("start DMX")
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal(e)
	}

	// Send RGB
	// 1: brightness/flash, 2: red, 3: blue, 4: green
	dmx.ChannelMap(34, 31, 32, 33)
	for {
		time.Sleep(500 * time.Millisecond)
		dmx.SendRGB(255, 255, 255, 255)
		log.Println("EXECUTED")
	}
	// Initial color.
	// dmx.SetChannel(34, 100)
	// dmx.SetChannel(32, 70)
	// dmx.SetChannel(33, 130)
	// dmx.SetChannel(31, 255)
	// dmx.Render()

	// for {

	// 	// Wait.
	// 	time.Sleep(90 * time.Millisecond)

	// 	dmx.SetChannel(34, 255) // Intensity
	// 	dmx.SetChannel(32, 255) // R
	// 	dmx.SetChannel(33, 255) // G
	// 	dmx.SetChannel(31, 255) // B
	// 	dmx.Render()

	// }
}
