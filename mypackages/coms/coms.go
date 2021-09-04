package coms

import (
	"fmt"
	"log"
	"net/url"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/yryz/ds18b20"
)

//Returns the temprature in Celsius
func Tempc() float64 {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	//fmt.Printf("sensor IDs: %v\n", sensors)

	for _, sensor := range sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			return t
		}
	}
	return 0
}

//Returns the temprature in Fahrenheit
func Tempf() float64 {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}
	//pass sensor in. If its blank return all temps. If it is not, return that temp.
	for _, sensor := range sensors {
		t, err := ds18b20.Temperature(sensor)
		if err == nil {
			f := t/(5/9) + 32
			return f
		}
	}
	return 0
}

//Returns a list of temp sensors (ds18b20's specifically)
func Tempsensors() float64 {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}

	fmt.Printf("sensor IDs: %v\n", sensors)
	return 0
}

//MQTT communications and routines

//place holders so as not to delete the import :-) taken from https://www.cloudmqtt.com/docs/go.html and https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang
func MQTTconnect(clientId string, uri *url.URL) mqtt.Client {
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

func MQTTcreateClientOptions(clientId string, uri *url.URL) *mqtt.ClientOptions {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s", uri.Host))
	opts.SetUsername(uri.User.Username())
	password, _ := uri.User.Password() 
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func MQTTlisten(uri *url.URL, topic string) {
	client := connect("sub", uri)
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}
