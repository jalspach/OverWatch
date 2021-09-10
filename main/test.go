package main

import (
	"fmt"
	"time"

	//"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
	"github.com/jalspach/OverWatch/mypackages/onewire"
	"github.com/jalspach/OverWatch/mypackages/reporting"
)

func main() {

	leds.Init()
	//leds.Error(500)
	//leds.SweepR2G(125)
	//leds.SweepG2R(125)
	//leds.Test()

	var cdeg float64 = onewire.Tempc()
	fmt.Printf("%fdeg deg C\n", cdeg)

	var fdeg float64 = onewire.Tempf()
	fmt.Printf("%Fdeg deg F\n", fdeg)

	onewire.Tempc()
	onewire.Tempf()

	leds.DisplayTemp()

	reporting.Noppers()
	reporting.Report()

	//
	//
	//

	time.Sleep(10 * time.Second)
	leds.Init()
}
