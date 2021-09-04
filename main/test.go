package main

import (
	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {
	leds.Error()
	leds.Test()
	coms.Tempc()
	leds.Reset()
	// Set run types
	// oneoff - test leds and continue once
	// continious - test leds once and then continue
	// service - skip LED test
	//
	// grab temp and network test results and report them to MQTT
	// set the LED's as apropriate for the network test

}
