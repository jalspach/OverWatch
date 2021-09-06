package main

import (
	"fmt"
	"time"

	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {

	leds.Init()
	//leds.Error(500)
	//leds.SweepR2G(125)
	//leds.SweepG2R(125)
	//leds.Test()

	var c float64 = coms.Tempc()
	fmt.Printf("%f deg C\n", c)

	var f float64 = coms.Tempf()
	fmt.Printf("%f deg F\n", f)

	coms.Tempc()
	coms.Tempf()

	leds.DisplayTemp()
	time.Sleep(10 * time.Second)
	leds.Init()

}
