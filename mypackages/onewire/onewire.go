package onewire

import (
	"fmt"

	"github.com/yryz/ds18b20"
)

//tewt
func Test() int {
	return 0
}

//Return list of ds18B20's
func Listb20s() int {
	sensors, err := ds18b20.Sensors()
	if err != nil {
		panic(err)
	}
	fmt.Printf("sensor IDs: %v\n", sensors)
	return 0
}

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
			f := t/(.556) + 32
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