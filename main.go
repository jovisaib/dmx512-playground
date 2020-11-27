package main

import (
	"log"
	"time"

	"github.com/akualab/dmx"
)

func main() {

	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal("ERROR: ", e)
	}

	for {

		time.Sleep(500 * time.Millisecond)

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

		dmx.SetChannel(31, 100)
		dmx.SetChannel(32, 70)
		dmx.SetChannel(33, 130)
		dmx.SetChannel(34, 180)

		dmx.SetChannel(36, 100)
		dmx.SetChannel(37, 70)
		dmx.SetChannel(38, 130)
		dmx.SetChannel(39, 180)

		dmx.SetChannel(41, 100)
		dmx.SetChannel(42, 70)
		dmx.SetChannel(43, 130)
		dmx.SetChannel(44, 180)

		if err := dmx.Render(); err != nil {
			log.Fatal("RENDER ERROR: ", err)
		}

		log.Println("RENDER!")
	}
}
