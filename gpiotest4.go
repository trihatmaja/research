package main

import (
	"fmt"
	"github.com/mrmorphic/hwio"
	"math"
	"strconv"
	"time"
)

func main() {
	m, e := hwio.GetModule("i2c")
	if e != nil {
		fmt.Printf("could not get i2c module: %s\n", e)
		return
	}
	i2c := m.(hwio.I2CModule)

	// Uncomment on Raspberry pi, which doesn't automatically enable i2c bus. BeagleBone does,
	// as the default device tree enables it.

	i2c.Enable()
	defer i2c.Disable()

	device := i2c.GetDevice(0x48)

	for {
		a, _ := device.Read(0x48, 1)
		s := fmt.Sprintf("%d", a[0])
		i, _ := strconv.ParseFloat(s, 64)
		fmt.Println(math.Abs(i-123) * 0.039)
		time.Sleep(1 * time.Second)
	}
}
