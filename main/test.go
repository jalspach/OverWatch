package main

import (
	"fmt"
	"time"

	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {

	// ToDo:
	// Upload data to MQTT (IP Address and temp)
	// Run network test(s)
	// Upload data to MQTT (test results)
	// set the LED's as apropriate (temp or network)
	//
	//Based on run type do the above as apropirate
	//
	// Set run types
	// oneoff - test leds, report temprature while doing nextork test? Threads?
	// continious - test leds, report temp while doing network test, loop over these two
	// service - skip LED test
	//
	// Maybe do this using threads if thats not too much extra work
	//


	leds.Init()
	//leds.Error(500)
	//leds.SweepR2G(125)
	//leds.SweepG2R(125)
	//leds.Test()

	var c float64 = coms.Tempc()
	fmt.Printf("%f\n", c)

	var f float64 = coms.Tempf()
	fmt.Printf("%f\n", f)

	coms.Tempc()
	coms.Tempf()

	leds.DisplayTemp()
	time.Sleep(10 * time.Second)
	leds.Init()

}
