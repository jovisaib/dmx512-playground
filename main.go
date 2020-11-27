package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/akualab/dmx"
)

func main() {
	r := rand.New(rand.NewSource(99))
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal("ERROR: ", e)
	}

	for {
		time.Sleep(1000 * time.Millisecond)

		// Set values for channels.
		dmx.SetChannel(1, byte(r.Intn(256)))
		dmx.SetChannel(2, byte(r.Intn(256)))
		dmx.SetChannel(3, byte(r.Intn(256)))
		dmx.SetChannel(4, byte(r.Intn(256)))

		dmx.SetChannel(5, byte(r.Intn(256)))
		dmx.SetChannel(6, byte(r.Intn(256)))
		dmx.SetChannel(7, byte(r.Intn(256)))
		dmx.SetChannel(8, byte(r.Intn(256)))

		dmx.SetChannel(9, byte(r.Intn(256)))
		dmx.SetChannel(10, byte(r.Intn(256)))
		dmx.SetChannel(11, byte(r.Intn(256)))
		dmx.SetChannel(12, byte(r.Intn(256)))

		dmx.SetChannel(13, byte(r.Intn(256)))
		dmx.SetChannel(14, byte(r.Intn(256)))
		dmx.SetChannel(15, byte(r.Intn(256)))
		dmx.SetChannel(16, byte(r.Intn(256)))

		dmx.SetChannel(26, byte(r.Intn(256)))
		dmx.SetChannel(27, byte(r.Intn(256)))
		dmx.SetChannel(28, byte(r.Intn(256)))
		dmx.SetChannel(29, byte(r.Intn(256)))

		dmx.SetChannel(31, byte(r.Intn(256)))
		dmx.SetChannel(32, byte(r.Intn(256)))
		dmx.SetChannel(33, byte(r.Intn(256)))
		dmx.SetChannel(34, byte(r.Intn(256)))

		dmx.SetChannel(36, byte(r.Intn(256)))
		dmx.SetChannel(37, byte(r.Intn(256)))
		dmx.SetChannel(38, byte(r.Intn(256)))
		dmx.SetChannel(39, byte(r.Intn(256)))

		dmx.SetChannel(41, byte(r.Intn(256)))
		dmx.SetChannel(42, byte(r.Intn(256)))
		dmx.SetChannel(43, byte(r.Intn(256)))
		dmx.SetChannel(44, byte(r.Intn(256)))

		if err := dmx.Render(); err != nil {
			log.Fatal("RENDER ERROR: ", err)
		}

		log.Println("RENDER!")
	}
}
