package main

import (
	"fmt"
	"os"
	"time"

	//"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
	"github.com/jalspach/OverWatch/mypackages/onewire"
	"github.com/jalspach/OverWatch/mypackages/reporting"
	"github.com/jalspach/OverWatch/mypackages/util"
)

// Publishes the deg F
func PublishTempF(client, basetopic string) int {
	temp := fmt.Sprintf("%.2f", onewire.Tempf())
	reporting.Publish(client, basetopic+"Temperature", temp, 0)
	return 0
}

func PublishIP(client, basetopic string) int {
	ipaddr, _ := util.GetInterfaceIpv4Addr("eth0")
	reporting.Publish(client, basetopic+"ip_address", ipaddr, 0)
	return 0
}

func PublishPortCheck(client, basetopic, targethost, targetport string) int {
	portstatus := util.PortCheckSimple(targethost, targetport)
	reporting.Publish(client, basetopic+"PortCheck", portstatus, 0)
	return 0
}

func main() {
	client, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	var targethost string = "slashdot.org"
	var targetport string = "443"
	var basetopic string = "home/frontroom/" + client + "/"
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
	PublishTempF(client, basetopic)
	PublishIP(client, basetopic)
	PublishPortCheck(client, basetopic, targethost, targetport)
	//util.PortCheck("www.slashdot.org", "80")

	// reporting.Publish(client, topic, value, qos)

	//
	//
	//

	time.Sleep(10 * time.Second)
	leds.Init()
}
