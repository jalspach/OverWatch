package reporting

import (
	"fmt"
	//import the Paho Go MQTT library
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//MQTT communications and routines

//place holders so as not to delete the import :-) taken from https://www.eclipse.org/paho/index.php?page=clients/golang/index.php https://www.cloudmqtt.com/docs/go.html and https://www.emqx.com/en/blog/how-to-use-mqtt-in-golang

//define a function for the default message handler

func init(client MQTT.Client, msg MQTT.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
