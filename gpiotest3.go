package main

import (
	// "fmt"
	"github.com/mrmorphic/hwio"
)

func main(){
	// hwio menggunakan bcm mode
	defer hwio.CloseAll()
	myPin0, _ := hwio.GetPinWithMode("gpio17", hwio.OUTPUT)
	for{
		e := hwio.Pulse(myPin0, hwio.HIGH, 100)
		if e!=nil{
			panic(e)
		}
	}
}
