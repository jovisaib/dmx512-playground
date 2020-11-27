package main

import (
	"log"
	"time"

	"github.com/jovisaib/dmx512-playground/dmx"
)

func main() {
	// r := rand.New(rand.NewSource(99))
	dmx, e := dmx.NewDMXConnection("/dev/ttyUSB0")
	if e != nil {
		log.Fatal("ERROR: ", e)
	}

	defer dmx.Close()

	for {
		time.Sleep(70 * time.Millisecond)

		dmx.SetChannel(26, 255)
		dmx.SetChannel(27, 255)
		dmx.SetChannel(28, 255)
		dmx.SetChannel(29, 255)

		if err := dmx.Render(); err != nil {
			log.Fatal("RENDER ERROR: ", err)
		}

		log.Println("RENDER 1!")

		dmx.SetChannel(31, 255)
		dmx.SetChannel(32, 255)
		dmx.SetChannel(33, 255)
		dmx.SetChannel(34, 255)

		if err := dmx.Render(); err != nil {
			log.Fatal("RENDER ERROR: ", err)
		}

		log.Println("RENDER 2!")

		dmx.SetChannel(36, 255)
		dmx.SetChannel(37, 255)
		dmx.SetChannel(38, 255)
		dmx.SetChannel(39, 255)

		if err := dmx.Render(); err != nil {
			log.Fatal("RENDER ERROR: ", err)
		}

		log.Println("RENDER 3!")

		dmx.SetChannel(41, 255)
		dmx.SetChannel(42, 255)
		dmx.SetChannel(43, 255)
		dmx.SetChannel(44, 255)

		if err := dmx.Render(); err != nil {
			log.Fatal("RENDER ERROR: ", err)
		}

		log.Println("RENDER 4!")
	}
}
