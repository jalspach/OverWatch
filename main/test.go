package main

import (
	"../mypackages/coms"
	"../mypackages/leds"
)

func main() {
	leds.Setstatus(0x0)
	leds.Error(5, 5)
	coms.Tempc()
	coms.Tempf()
	coms.Tempsensors()
}
