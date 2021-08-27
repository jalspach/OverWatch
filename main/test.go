package main

import (
	"log"
	"net/url"
	"os"
	"time"

	"github.com/jalspach/OverWatch/mypackages/coms"
	"github.com/jalspach/OverWatch/mypackages/leds"
)

func main() {
	leds.Test()
	leds.Error(3, 3)
	leds.Test()
	coms.Tempc()

	uri, err := url.Parse(os.Getenv("CLOUDMQTT_URL"))
	if err != nil {
		log.Fatal(err)
	}
	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "test"
	}

	go listen(uri, topic)

	client := coms.MQTTconnect("pub", uri)
	timer := time.NewTicker(1 * time.Second)
	for t := range timer.C {
		client.Publish(topic, 0, false, t.String())
	}

}
