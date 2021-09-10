package mqtt

import (
	"log"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

//MQTT communications and routines

//place holders so as not to delete the import :-) taken from https://www.cloudmqtt.com/docs/go.html and https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang

//MQTTInit
func Init(clientId string, uri *url.URL) mqtt.Client {
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

/*
func createClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", ebcd79b43b7f44a2a105831f0070400a.s1.eu.hivemq.cloud))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password()
	opts.SetPassword(Gr33tings)
	opts.SetClientID(devices)
	return opts
}

/*
func MQTTlisten(uri *url.URL, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}
*/
