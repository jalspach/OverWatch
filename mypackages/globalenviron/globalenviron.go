package globalenviron

import "github.com/garethpaul/purpleair-go"

// Series of functions to read from Purple Air. These only pull the first data point. Need to fix that most likely.
// https://github.com/garethpaul/purpleair-go/blob/f06144ce241d44212feb26d73136e9f89232051b/results.go

//Returns the label based on the sensor ID provided
func GetLabel(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Label
}

//Returns the label based on the sensor ID provided
func GetLocationType(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].DEVICELOCATIONTYPE
}

//Returns the label based on the sensor ID provided
func GetLat(sensor string) float64 {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Lat
}

//Returns the label based on the sensor ID provided
func GetLon(sensor string) float64 {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Lon
}

//Returns the label based on the sensor ID provided
func GetPM25Value(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Label
}

//Returns the label based on the sensor ID provided
func GetLastSeen(sensor string) int {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].LastSeen
}

//Returns the label based on the sensor ID provided
func GetType(sensor string) string {
	client := purpleair.NewClient()
	s := client.Sensor(sensor)
	return s.Results[0].Type
}
