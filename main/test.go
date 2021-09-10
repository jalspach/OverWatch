package main

import (
	"fmt"
	"time"

	//"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
	"github.com/jalspach/OverWatch/mypackages/mqtt"
	"github.com/jalspach/OverWatch/mypackages/onewire"
)

func main() {

	leds.Init()
	//leds.Error(500)
	//leds.SweepR2G(125)
	//leds.SweepG2R(125)
	//leds.Test()

	var c float64 = onewire.Tempc()
	fmt.Printf("%f deg C\n", c)

	var f float64 = onewire.Tempf()
	fmt.Printf("%f deg F\n", f)

	onewire.Tempc()
	onewire.Tempf()

	leds.DisplayTemp()
	time.Sleep(10 * time.Second)
	leds.Init()

	mqtt.Init()

}
