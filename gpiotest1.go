package main

import (
	"github.com/mrmorphic/hwio"
)

func main() {
	// hwio menggunakan bcm mode
	triggerPin, err := hwio.GetPinWithMode("gpio17", hwio.OUTPUT)
	zeroPin, err := hwio.GetPinWithMode("gpio18", hwio.INPUT)
	if err != nil {
		panic(err)
	}
	for {
		val, _ := hwio.DigitalRead(zeroPin)
		if val == hwio.HIGH {
			hwio.DigitalWrite(triggerPin, hwio.HIGH)
			hwio.DelayMicroseconds(100)
			hwio.DigitalWrite(triggerPin, hwio.LOW)
			hwio.DelayMicroseconds(9900)
		}
	}
}
