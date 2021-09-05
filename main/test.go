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
	fmt.Print(coms.Tempc())
	fmt.Print(coms.Tempf())

	//leds.Init()

	// Set run types
	// oneoff - test leds, report temprature while doing nextork test? Threads?
	// continious - test leds, report temp while doing network test, loop over these two
	// service - skip LED test
	//
	// grab temp and network test results and report them to MQTT
	// set the LED's as apropriate for the network test

}
