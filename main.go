package main

import (
	"log"

	"github.com/oliread/usbdmx"
	"github.com/oliread/usbdmx/ft232"
)

func main() {
	vid := 1
	pid := 1
	outputInterfaceID := 1
	inputInterfaceID := 1
	debugLevel := 1

	config := usbdmx.NewConfig(uint16(vid), uint16(pid), outputInterfaceID, inputInterfaceID, debugLevel)
	config.GetUSBContext()

	controller := ft232.NewDMXController(config)
	if err := controller.Connect(); err != nil {
		log.Fatal(err)
	}

	controller.SetChannel(1, 255)
	controller.SetChannel(2, 255)
	controller.SetChannel(3, 255)
	controller.SetChannel(4, 255)

	if err := controller.Render(); err != nil {
		log.Fatal(err)
	}
}
