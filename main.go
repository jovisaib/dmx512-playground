package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/jovisaib/dmx512-playground/dmx"
)

func main() {
	r := rand.New(rand.NewSource(99))
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal("ERROR: ", e)
	}

	for {
		time.Sleep(1000 * time.Millisecond)

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
