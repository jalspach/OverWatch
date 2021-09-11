package main

import (
	"fmt"
	"time"

	//"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
	"github.com/jalspach/OverWatch/mypackages/onewire"
	"github.com/jalspach/OverWatch/mypackages/reporting"
)

// Publishes the deg F
func Publishtempf(client string) int {
	temp := fmt.Sprintf("%.2f", onewire.Tempf())
	reporting.Publish(client, "Temperature", temp, 0)

	return 0
}

func main() {
	var client string = "Hostname"
	//	var topic string = "topic"
	//	var qos byte = 0
	//	var value string

	leds.Init()
	//leds.Error(500)
	//leds.SweepR2G(125)
	//leds.SweepG2R(125)
	//leds.Test()

	var cdeg float64 = onewire.Tempc()
	fmt.Printf("%fdeg deg C\n", cdeg)

	var fdeg float64 = onewire.Tempf()
	fmt.Printf("%Fdeg deg F\n", fdeg)

	//	onewire.Tempc()
	//	temp := fmt.Sprintf("%v", onewire.Tempf())
	//	value = temp

	leds.DisplayTemp()

	//seems like I should build a struct that gets populated and passed to report
	Publishtempf(client)

	// reporting.Publish(client, topic, value, qos)

	//
	//
	//

	time.Sleep(10 * time.Second)
	leds.Init()
}
