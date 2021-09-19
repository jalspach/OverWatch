package globalenviron

import "github.com/garethpaul/purpleair-go"

// Series of functions to read from Purple Air

//Returns the label based on the sensor ID provided
func GetLabel(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Label
}
