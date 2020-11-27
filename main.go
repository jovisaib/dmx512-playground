package main

import (
	"github.com/akualab/dmx"
	"log"
	"math/rand"
	"time"
)

func main() {

	r := rand.New(rand.NewSource(99))
	log.Printf("start DMX")
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal(e)
	}

	// Send RGB
	// 1: brightness/flash, 2: red, 3: blue, 4: green
	dmx.ChannelMap(34, 31, 33, 32)
	dmx.SendRGB(130, 100, 100, 100)
	time.Sleep(10 * time.Second)

	// Initial color.
	dmx.SetChannel(34, 100)
	dmx.SetChannel(32, 70)
	dmx.SetChannel(33, 130)
	dmx.SetChannel(31, 255)
	dmx.Render()

	for {

		// Wait.
		time.Sleep(90 * time.Millisecond)

		dmx.SetChannel(34, 255)               // Intensity
		dmx.SetChannel(32, byte(r.Intn(256))) // R
		dmx.SetChannel(33, byte(r.Intn(256))) // G
		dmx.SetChannel(31, byte(r.Intn(256))) // B
		dmx.Render()

	}
}
