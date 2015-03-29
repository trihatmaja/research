package main

import (
	"github.com/mrmorphic/hwio"
)

func main() {
	// hwio menggunakan bcm mode
	defer hwio.CloseAll()
	myPin0, err := hwio.GetPinWithMode("gpio17", hwio.OUTPUT)
	myPin1, err := hwio.GetPinWithMode("gpio18", hwio.INPUT)
	if err != nil {
		panic(err)
	}

	for {
		value, _ := hwio.DigitalRead(myPin1)
		if value == hwio.HIGH {
			hwio.DigitalWrite(myPin0, hwio.HIGH)
		} else {
			hwio.DigitalWrite(myPin0, hwio.LOW)
		}
	}
}
