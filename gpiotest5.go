package main

import (
	"os"
)

func main() {
	const fifo = "/dev/pi-blaster"
	f, err := os.Create(fifo)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	_, err = f.WriteString("17=0\n")
	if err != nil {
		panic(err)
	}
	f.Sync()
}
