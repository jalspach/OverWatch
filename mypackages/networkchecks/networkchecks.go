package networkchecks

import (
	"fmt"
	//import the Paho Go MQTT library

	"github.com/showwin/speedtest-go/speedtest"
)

//Accept destination as a fqdn or server number (Include SCOE with any request so as to provide a baseline. Maybe limit to a single destination with SCOE as the default).
//Return the data as a slice
func netcheck() {
	user, _ := speedtest.FetchUserInfo()

	serverList, _ := speedtest.FetchServerList(user)
	targets, _ := serverList.FindServer([]int{})

	//What else is available? Packet Loss? Jitter?
	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(false)
		s.UploadTest(false)

		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
	}
}
