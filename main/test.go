package main

import (
	"fmt"
	"math"
	"os"
	"time"

	//"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
	"github.com/jalspach/OverWatch/mypackages/onewire"
	"github.com/jalspach/OverWatch/mypackages/reporting"
	"github.com/jalspach/OverWatch/mypackages/util"
)

// Publishes the deg F to MQTT (checks 3 times and send the middle)
func PublishTempF(client string, basetopic string, qos byte) int {
	t1 := onewire.Tempf()
	t2 := onewire.Tempf()
	t3 := onewire.Tempf()
	temp := math.Max(math.Min(t1, t2), math.Min(math.Max(t1, t2), t3))
	reporting.Publish(client, basetopic+"Temperature", fmt.Sprintf("%.2f", temp), qos)
	return 0
}

//Figureout our IPv4 address and publish it to MQTT
func PublishIP(client string, basetopic string, qos byte) int {
	ipaddr, _ := util.GetInterfaceIpv4Addr("eth0")
	reporting.Publish(client, basetopic+"IPAddress", ipaddr, qos)
	return 0
}

//test to a port at a server and publish that to MQTT. Sends txt to represnet status
func PublishSimplePortCheckTxt(client string, basetopic string, targethost string, targetport string, qos byte) int {
	portstatus := util.PortCheckSimpleTxt(targethost, targetport)
	reporting.Publish(client, basetopic+"PortCheck", portstatus, qos)
	return 0
}

//test to a port at a server and publish that to MQTT. Sends 0 or 1 to rep status
func PublishSimplePortCheckBool(client string, basetopic string, targethost string, targetport string, qos byte) int {
	portstatus := util.PortCheckSimpleBool(targethost, targetport)
	reporting.Publish(client, basetopic+"HostPortUp", portstatus, qos)
	return 0
}

//Display Temprature on LED's
func DisplayTemp() int {
	//coms.Listb20s()
	var f float64 = onewire.Tempf()
	//fmt.Printf("%f\n", f)

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
	return 0
}

func main() {
	//setups
	client, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	var targethost string = "slashdot.org"
	var targetport string = "443"
	var basetopic string = "home/frontroom/" + client + "/"
	var qos byte = 0

	//startin
	leds.Init()
	//leds.Error(500)
	//leds.SweepR2G(125)
	//leds.SweepG2R(125)
	//leds.Test()

	var cdeg float64 = onewire.Tempc()
	fmt.Printf("%fdeg deg C\n", cdeg)

	var fdeg float64 = onewire.Tempf()
	fmt.Printf("%Fdeg deg F\n", fdeg)

	util.CheckAQI("15471")
	//seems like I should build a struct that gets populated and passed to publish events

	go leds.SweepG2R(65)
	PublishTempF(client, basetopic, qos)
	PublishIP(client, basetopic, qos)
	PublishSimplePortCheckTxt(client, basetopic, targethost, targetport, qos)
	PublishSimplePortCheckBool(client, basetopic, targethost, targetport, qos)
	leds.Setstatus(0)
	DisplayTemp()

	time.Sleep(10 * time.Second)
	leds.Init()

}
