package main

import (
	"fmt"

	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {
	leds.Init()
	//leds.Error(500)
	//leds.SweepR2G(125)
	//leds.SweepG2R(125)
	//leds.Test()
	coms.Listb20s()
	var c float64 = coms.Tempc()
	fmt.Printf("%f\n", c)

	f := c/(.556) + 32
	fmt.Printf("%f\n", f)

	//leds.Init()

	// Set run types
	// oneoff - test leds, report temprature while doing nextork test? Threads?
	// continious - test leds, report temp while doing network test, loop over these two
	// service - skip LED test
	//
	// grab temp and network test results and report them to MQTT
	// set the LED's as apropriate for the network test

	switch {
	case f > 100:
		leds.Setstatus(0x_1)
	case f > 90:
		leds.Setstatus(0x_3)
	case f > 84:
		leds.Setstatus(0x_6)
	case f > 75:
		leds.Setstatus(0x_4)
	case f > 68:
		leds.Setstatus(0x_6)
	case f > 60:
		leds.Setstatus(0x_3)
	default:
		leds.Setstatus(0x_1)
	}
}
